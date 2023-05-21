package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Customer struct {
	ID        uint      `gorm:"primarykey"`
	FirstName string    `gorm:"type:varchar(100);not null"`
	LastName  string    `gorm:"type:varchar(100);not null"`
	BirthDate time.Time `gorm:"not null"`
	Gender    string    `gorm:"type:varchar(6);not null"`
	Email     string    `gorm:"type:varchar(100);not null;unique"`
	Address   string    `gorm:"type:varchar(200)"`
}

func main() {
	dsn := "host=localhost user=postgres password=password dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	performMigration(db)

	customerRepo := NewCustomerRepository(db)
	customerService := NewCustomerService(customerRepo)
	customerController := NewCustomerController(customerService)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", customerController.GetCustomers)
	r.GET("/edit/:id", customerController.GetCustomerByID)
	r.GET("/create", customerController.ShowCreateForm)
	r.POST("/create", customerController.CreateCustomer)
	r.POST("/edit/:id", customerController.UpdateCustomer)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
func performMigration(db *gorm.DB) {
	err := db.AutoMigrate(&Customer{})
	if err != nil {
		log.Fatal(err)
	}
}
