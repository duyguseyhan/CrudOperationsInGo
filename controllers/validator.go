package controllers

import (
	"errors"
	"regexp"
	"strings"
	"time"

	m "github.com/duyguseyhan/crudoperationsingo/models"
)

func validateCustomer(customer m.Customer) error {
	if strings.TrimSpace(customer.FirstName) == "" {
		return errors.New("first name is required")
	}

	if len(strings.TrimSpace(customer.FirstName)) > 100 {
		return errors.New("first name length exceeds the limit")
	}

	if strings.TrimSpace(customer.LastName) == "" {
		return errors.New("last name is required")
	}

	if len(strings.TrimSpace(customer.LastName)) > 100 {
		return errors.New("last name length exceeds the limit")
	}

	now := time.Now()
	minAge := 18
	maxAge := 60
	age := now.Year() - customer.BirthDate.Year()
	if customer.BirthDate.After(now.AddDate(-age, 0, 0)) {
		age--
	}

	if age < minAge || age > maxAge {
		return errors.New("age should be between 18 and 60")
	}

	if customer.Gender != "Male" && customer.Gender != "Female" {
		return errors.New("invalid gender")
	}

	if strings.TrimSpace(customer.Email) == "" {
		return errors.New("email is required")
	}

	if !validateEmail(strings.TrimSpace(customer.Email)) {
		return errors.New("invalid email address")
	}

	if len(strings.TrimSpace(customer.Address)) > 200 {
		return errors.New("address length exceeds the limit")
	}

	return nil
}

func validateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
