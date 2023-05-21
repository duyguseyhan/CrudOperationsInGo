package services

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	model "github.com/duyguseyhan/crudoperationsingo/models"
)

type MockCustomerRepository struct {
	mock.Mock
}

// GetCustomerByID implements repositories.CustomerRepository
func (m *MockCustomerRepository) GetCustomerByID(id string) (*model.Customer, error) {
	args := m.Called(id)
	return args.Get(0).(*model.Customer), args.Error(1)
}

// GetCustomers implements repositories.CustomerRepository
func (m *MockCustomerRepository) GetCustomers() ([]model.Customer, error) {
	args := m.Called()
	customers := make([]model.Customer, 0)
	if result, ok := args.Get(0).([]model.Customer); ok {
		customers = result
	}
	return customers, args.Error(1)
}

// SearchCustomers implements repositories.CustomerRepository
func (m *MockCustomerRepository) SearchCustomers(query string) ([]model.Customer, error) {
	args := m.Called(query)
	return args.Get(0).([]model.Customer), args.Error(1)
}

// SortCustomers implements repositories.CustomerRepository
func (m *MockCustomerRepository) SortCustomers(column string, descending bool) ([]model.Customer, error) {
	args := m.Called(column, descending)
	return args.Get(0).([]model.Customer), args.Error(1)
}

func (m *MockCustomerRepository) CreateCustomer(customer model.Customer) error {
	args := m.Called(customer)
	return args.Error(0)
}

func (m *MockCustomerRepository) UpdateCustomer(id string, customer model.Customer) error {
	args := m.Called(id, customer)
	return args.Error(0)
}

func TestUpdateCustomer(t *testing.T) {
	repo := new(MockCustomerRepository)
	service := NewCustomerService(repo)

	// Define the input parameters
	customerID := "1"
	newCustomer := model.Customer{
		FirstName: "Updated",
		LastName:  "Customer",
	}

	// Define the expected error (if any)
	expectedError := errors.New("database error")

	// Mock the repository method
	repo.On("UpdateCustomer", customerID, newCustomer).Return(expectedError)

	// Call the method being tested
	err := service.UpdateCustomer(customerID, newCustomer)

	// Assert the result
	assert.Equal(t, expectedError, err)

	// Assert that the repository method was called
	repo.AssertCalled(t, "UpdateCustomer", customerID, newCustomer)
}
func TestCreateCustomer(t *testing.T) {
	repo := new(MockCustomerRepository)
	service := NewCustomerService(repo)

	// Define the input parameter
	newCustomer := model.Customer{
		FirstName: "John",
		LastName:  "Doe",
	}

	// Mock the repository method
	repo.On("CreateCustomer", newCustomer).Return(nil)

	// Call the method being tested
	err := service.CreateCustomer(newCustomer)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the repository method was called
	repo.AssertCalled(t, "CreateCustomer", newCustomer)
}
func TestGetCustomerByID(t *testing.T) {
	repo := new(MockCustomerRepository)
	service := NewCustomerService(repo)

	// Define the input parameter
	customerID := "1"

	// Define the expected result
	expectedCustomer := &model.Customer{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
	}
	repo.On("GetCustomerByID", customerID).Return(expectedCustomer, nil)

	// Call the method being tested
	customer, err := service.GetCustomerByID(customerID)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomer, customer)

	// Assert that the repository method was called
	repo.AssertCalled(t, "GetCustomerByID", customerID)
}

func TestSearchCustomers(t *testing.T) {
	repo := new(MockCustomerRepository)
	service := NewCustomerService(repo)

	// Define the input parameter
	query := "John"

	// Define the expected result
	expectedCustomers := []model.Customer{
		{ID: 1, FirstName: "John", LastName: "Doe"},
		{ID: 2, FirstName: "Johnny", LastName: "Smith"},
	}
	repo.On("SearchCustomers", query).Return(expectedCustomers, nil)

	// Call the method being tested
	customers, err := service.SearchCustomers(query)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomers, customers)

	// Assert that the repository method was called
	repo.AssertCalled(t, "SearchCustomers", query)
}

func TestSortCustomers(t *testing.T) {
	repo := new(MockCustomerRepository)
	service := NewCustomerService(repo)

	// Define the input parameters
	column := "FirstName"
	descending := false

	// Define the expected result
	expectedCustomers := []model.Customer{
		{ID: 2, FirstName: "Jane", LastName: "Smith"},
		{ID: 1, FirstName: "John", LastName: "Doe"},
	}
	repo.On("SortCustomers", column, descending).Return(expectedCustomers, nil)

	// Call the method being tested
	customers, err := service.SortCustomers(column, descending)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomers, customers)

	// Assert that the repository method was called
	repo.AssertCalled(t, "SortCustomers", column, descending)
}

func TestGetCustomers(t *testing.T) {
	repo := new(MockCustomerRepository)
	service := NewCustomerService(repo)

	// Define the expected result
	expectedCustomers := []model.Customer{
		{ID: 1, FirstName: "John", LastName: "Doe"},
		{ID: 2, FirstName: "Jane", LastName: "Smith"},
		{ID: 3, FirstName: "Alice", LastName: "Johnson"},
	}
	repo.On("GetCustomers").Return(expectedCustomers, nil)

	// Call the method being tested
	customers, err := service.GetCustomers()

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedCustomers, customers)

	// Assert that the repository method was called
	repo.AssertCalled(t, "GetCustomers")
}
