package main

type CustomerService interface {
	GetCustomers() ([]Customer, error)
	CreateCustomer(customer Customer) error
	UpdateCustomer(id string, customer Customer) error
	GetCustomerByID(id string) (*Customer, error)
	SearchCustomers(query string) ([]Customer, error)
	SortCustomers(column string, descending bool) ([]Customer, error)
}

type customerService struct {
	repo CustomerRepository
}

func NewCustomerService(repo CustomerRepository) CustomerService {
	return &customerService{
		repo: repo,
	}
}

func (s *customerService) GetCustomers() ([]Customer, error) {
	return s.repo.GetCustomers()
}

func (s *customerService) CreateCustomer(customer Customer) error {
	return s.repo.CreateCustomer(customer)
}

func (s *customerService) UpdateCustomer(id string, customer Customer) error {
	return s.repo.UpdateCustomer(id, customer)
}

func (s *customerService) GetCustomerByID(id string) (*Customer, error) {
	return s.repo.GetCustomerByID(id)
}

func (s *customerService) SearchCustomers(query string) ([]Customer, error) {
	return s.repo.SearchCustomers(query)
}

func (s *customerService) SortCustomers(column string, descending bool) ([]Customer, error) {
	return s.repo.SortCustomers(column, descending)
}
