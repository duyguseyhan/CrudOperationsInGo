// validator.go

package main

import (
	"errors"
	"regexp"
	"time"
)

func validateCustomer(customer Customer) error {
	if customer.FirstName == "" {
		return errors.New("first name is required")
	}

	if customer.LastName == "" {
		return errors.New("last name is required")
	}

	if customer.Gender != "Male" && customer.Gender != "Female" {
		return errors.New("invalid gender")
	}

	if !validateEmail(customer.Email) {
		return errors.New("invalid email address")
	}

	_, err := time.Parse("2006-01-02", customer.BirthDate)
	if err != nil {
		return errors.New("invalid birth date")
	}

	if len(customer.Address) > 100 {
		return errors.New("address length exceeds the limit")
	}

	return nil
}

func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
