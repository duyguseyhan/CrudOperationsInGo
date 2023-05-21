package services

import (
	m "github.com/duyguseyhan/crudoperationsingo/models"
	rep "github.com/duyguseyhan/crudoperationsingo/repositories"
)

type CustomerService interface {
	GetCustomers() ([]m.Customer, error)
	CreateCustomer(customer m.Customer) error
	UpdateCustomer(id string, customer m.Customer) error
	GetCustomerByID(id string) (*m.Customer, error)
	SearchCustomers(query string) ([]m.Customer, error)
	SortCustomers(column string, descending bool) ([]m.Customer, error)
}

type customerService struct {
	repo rep.CustomerRepository
}

func NewCustomerService(repo rep.CustomerRepository) CustomerService {
	return &customerService{
		repo: repo,
	}
}

func (s *customerService) GetCustomers() ([]m.Customer, error) {
	return s.repo.GetCustomers()
}

func (s *customerService) CreateCustomer(customer m.Customer) error {
	return s.repo.CreateCustomer(customer)
}

func (s *customerService) UpdateCustomer(id string, customer m.Customer) error {
	return s.repo.UpdateCustomer(id, customer)
}

func (s *customerService) GetCustomerByID(id string) (*m.Customer, error) {
	return s.repo.GetCustomerByID(id)
}

func (s *customerService) SearchCustomers(query string) ([]m.Customer, error) {
	return s.repo.SearchCustomers(query)
}

func (s *customerService) SortCustomers(column string, descending bool) ([]m.Customer, error) {
	return s.repo.SortCustomers(column, descending)
}
