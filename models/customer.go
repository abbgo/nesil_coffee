package models

import (
	"context"
	"errors"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
)

type Customer struct {
	ID       string `json:"id,omitempty"`
	Login    string `json:"login" binding:"required"`
	Mail     string `json:"mail" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ValidateRegisterCustomer(customer Customer) error {
	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if !helpers.IsEmailValid(customer.Mail) {
		return errors.New("invalid mail address")
	}

	var login string
	db.QueryRow(context.Background(), "SELECT login FROM customers WHERE login = $1", customer.Login).Scan(&login)
	if login != "" {
		return errors.New("this customer already exists")
	}

	return nil
}
