package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController interface {
	GetCustomers(c *gin.Context)
	CreateCustomer(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	GetCustomerByID(c *gin.Context)
	ShowCreateForm(c *gin.Context)
}

type customerController struct {
	service CustomerService
}

func NewCustomerController(service CustomerService) CustomerController {
	return &customerController{
		service: service,
	}
}

func (c *customerController) GetCustomers(ctx *gin.Context) {
	searchQuery := ctx.Query("search")
	sortColumn := ctx.Query("sort")
	descending := ctx.Query("desc") == "true"
	var customers []Customer
	var err error

	if searchQuery != "" {
		customers, err = c.service.SearchCustomers(searchQuery)
	} else if sortColumn != "" {
		customers, err = c.service.SortCustomers(sortColumn, descending)
	} else {
		customers, err = c.service.GetCustomers()
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"customers": customers,
	})
}

func (c *customerController) CreateCustomer(ctx *gin.Context) {
	var customer Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.CreateCustomer(customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *customerController) UpdateCustomer(ctx *gin.Context) {
	id := ctx.Param("id")

	var customer Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.service.UpdateCustomer(id, customer); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update customer"})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *customerController) GetCustomerByID(ctx *gin.Context) {
	id := ctx.Param("id")
	customer, err := c.service.GetCustomerByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.HTML(http.StatusOK, "edit.html", customer)
}
func (c *customerController) ShowCreateForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "create.html", nil)
}
