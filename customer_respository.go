package main

import (
	"errors"
	"time"

	"gorm.io/gorm"
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

type CustomerRepository interface {
	GetCustomers() ([]Customer, error)
	CreateCustomer(customer Customer) error
	UpdateCustomer(id string, customer Customer) error
	GetCustomerByID(id string) (*Customer, error)
	SearchCustomers(query string) ([]Customer, error)
	SortCustomers(column string, descending bool) ([]Customer, error)
}

type customerRepository struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) GetCustomers() ([]Customer, error) {
	var customers []Customer
	result := r.db.Order("first_name ASC").Find(&customers)
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

func (r *customerRepository) CreateCustomer(customer Customer) error {
	result := r.db.Create(&customer)
	return result.Error
}

func (r *customerRepository) UpdateCustomer(id string, customer Customer) error {
	result := r.db.Model(&Customer{}).Where("id = ?", id).Updates(customer)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("edit conflict: customer has been modified by another user")
	}
	return nil
}

func (r *customerRepository) GetCustomerByID(id string) (*Customer, error) {
	var customer Customer
	result := r.db.First(&customer, id)
	if result.Error != nil {
		return nil, result.Error
	}

	birthDate, _ := time.Parse("2006-01-02T15:04:05Z", customer.BirthDate)
	customer.BirthDate = birthDate.Format("2006-01-02")
	return &customer, nil
}

func (r *customerRepository) SearchCustomers(query string) ([]Customer, error) {
	var customers []Customer
	result := r.db.Where("first_name ILIKE ? OR last_name ILIKE ?", "%"+query+"%", "%"+query+"%").Find(&customers)
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

func (r *customerRepository) SortCustomers(column string, descending bool) ([]Customer, error) {
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
	result := r.db.Order(order).Find(&customers)
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
