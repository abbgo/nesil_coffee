package models

import (
	"context"
	"errors"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
)

type Customer struct {
	ID       string `json:"id,omitempty"`
	Login    string `json:"login,omitempty" binding:"required"`
	Mail     string `json:"mail,omitempty" binding:"required"`
	Password string `json:"password,omitempty" binding:"required"`
}

type LoginCustomer struct {
	Login    string `json:"login" binding:"required"`
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

func ValidateLoginCustomer(customer LoginCustomer) (Customer, error) {
	db, err := config.ConnDB()
	if err != nil {
		return Customer{}, err
	}
	defer db.Close()

	var returnCustomer Customer
	db.QueryRow(context.Background(), "SELECT id,login,mail,password FROM customers WHERE login = $1", customer.Login).
		Scan(&returnCustomer.ID, &returnCustomer.Login, &returnCustomer.Mail, &returnCustomer.Password)
	// eger request - den gelen telefon belgili admin database - de yok bolsa onda error response edilyar
	if returnCustomer.ID == "" {
		return Customer{}, errors.New("customer not found")
	}

	// eger admin bar bolsa onda paroly dogry yazypdyrmy yazmandyrmy sol barlanyar
	credentialError := helpers.CheckPassword(customer.Password, returnCustomer.Password)
	if credentialError != nil {
		return Customer{}, errors.New("invalid credentials")
	}

	returnCustomer.Password = ""

	return returnCustomer, nil
}
