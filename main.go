package main

import (
	"errors"
	"log"
	"net/http"
	"regexp"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Customer struct {
	ID        int
	FirstName string
	LastName  string
	BirthDate string
	Gender    string
	Email     string
	Address   string
}

func main() {
	dsn := "host=localhost user=postgres password=password dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", func(c *gin.Context) {
		searchQuery := c.Query("search")
		sortColumn := c.Query("sort")
		descending := c.Query("desc") == "true"
		var customers []Customer
		var err error

		if searchQuery != "" {
			customers, err = searchCustomers(db, searchQuery)
		} else if sortColumn != "" {
			customers, err = sortCustomers(db, sortColumn, descending)
		} else {
			customers, err = getCustomers(db)
		}

		if err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "index.html", gin.H{
			"customers": customers,
		})
	})

	r.GET("/edit/:id", func(c *gin.Context) {
		id := c.Param("id")
		customer, err := getCustomerByID(db, id)
		if err != nil {
			log.Fatal(err)
		}

		c.HTML(http.StatusOK, "edit.html", customer)
	})

	r.GET("/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", nil)
	})

	r.POST("/create", func(c *gin.Context) {
		var customer Customer
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !validateEmail(customer.Email) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
			return
		}

		if err := createCustomer(db, customer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
			return
		}

		c.Status(http.StatusOK)
	})

	r.POST("/edit/:id", func(c *gin.Context) {
		id := c.Param("id")

		var customer Customer
		if err := c.ShouldBindJSON(&customer); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if !validateEmail(customer.Email) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
			return
		}

		if err := updateCustomer(db, id, customer); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
			return
		}

		c.Status(http.StatusOK)
	})

	r.Run(":8080")
}

func getCustomers(db *gorm.DB) ([]Customer, error) {
	var customers []Customer
	result := db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	for i := range customers {
		birthDate, err := time.Parse("2006-01-02T15:04:05Z", customers[i].BirthDate)
		if err != nil {
			return nil, err
		}
		customers[i].BirthDate = birthDate.Format("02.01.2006")
	}
	return customers, nil
}

func createCustomer(db *gorm.DB, customer Customer) error {
	birthDate, _ := time.Parse("2006-01-02", customer.BirthDate)
	customer.BirthDate = birthDate.Format("2006-01-02")

	result := db.Create(&customer)
	return result.Error
}

func validateEmail(email string) bool {
	// Simple email validation using regular expression
	// You can use a more comprehensive library for email validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func getCustomerByID(db *gorm.DB, id string) (*Customer, error) {
	var customer Customer
	result := db.First(&customer, id)
	if result.Error != nil {
		return nil, result.Error
	}

	birthDate, _ := time.Parse("2006-01-02T15:04:05Z", customer.BirthDate)
	customer.BirthDate = birthDate.Format("2006-01-02")
	return &customer, nil
}

func updateCustomer(db *gorm.DB, id string, customer Customer) error {
	existingCustomer, err := getCustomerByID(db, id)
	if err != nil {
		return err
	}
	if existingCustomer == nil {
		return errors.New("customer not found")
	}

	if err := validateUpdatedCustomer(customer); err != nil {
		return err
	}

	birthDate, _ := time.Parse("2006-01-02", customer.BirthDate)
	customer.BirthDate = birthDate.Format("2006-01-02")

	result := db.Model(&existingCustomer).Updates(customer)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("edit conflict: customer has been modified by another user")
	}

	return nil
}

// Validate the updated customer data
func validateUpdatedCustomer(customer Customer) error {
	// Perform your validation checks here
	if customer.FirstName == "" {
		return errors.New("first name is required")
	}

	if customer.LastName == "" {
		return errors.New("last name is required")
	}

	// Add more validation checks as needed

	return nil
}

func searchCustomers(db *gorm.DB, query string) ([]Customer, error) {
	var customers []Customer
	result := db.Where("first_name ILIKE ? OR last_name ILIKE ?", "%"+query+"%", "%"+query+"%").Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	for i := range customers {
		birthDate, err := time.Parse("2006-01-02T15:04:05Z", customers[i].BirthDate)
		if err != nil {
			return nil, err
		}
		customers[i].BirthDate = birthDate.Format("02.01.2006")
	}
	return customers, nil
}
func sortCustomers(db *gorm.DB, column string, descending bool) ([]Customer, error) {
	var columnName string

	switch column {
	case "FirstName":
		columnName = "first_name"
	case "LastName":
		columnName = "last_name"
	case "BirthDate":
		columnName = "birth_date"
	case "Gender":
		columnName = "gender"
	case "Email":
		columnName = "email"
	case "Address":
		columnName = "address"
	default:
		return nil, errors.New("invalid sort column")
	}

	var order string
	if descending {
		order = columnName + " DESC"
	} else {
		order = columnName + " ASC"
	}

	var customers []Customer
	result := db.Order(order).Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}

	for i := range customers {
		birthDate, err := time.Parse("2006-01-02T15:04:05Z", customers[i].BirthDate)
		if err != nil {
			return nil, err
		}
		customers[i].BirthDate = birthDate.Format("02.01.2006")
	}

	return customers, nil
}
