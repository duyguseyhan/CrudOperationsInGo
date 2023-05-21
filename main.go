package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":5432)/" + os.Getenv("DB_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

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
