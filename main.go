package main

import (
	"log"

	m "github.com/duyguseyhan/crudoperationsingo/models"
	r "github.com/duyguseyhan/crudoperationsingo/repositories"
	s "github.com/duyguseyhan/crudoperationsingo/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=password dbname=mydb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	performMigration(db)

	customerRepo := r.NewCustomerRepository(db)
	customerService := s.NewCustomerService(customerRepo)
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
	err := db.AutoMigrate(&m.Customer{})
	if err != nil {
		log.Fatal(err)
	}
}
