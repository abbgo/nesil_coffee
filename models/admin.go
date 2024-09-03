package models

import (
	"context"
	"errors"
	"nesil_coffe/config"
)

type Admin struct {
	ID       string `json:"id,omitempty"`
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func ValidateRegisterAdmin(admin Admin) error {
	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	var login string
	db.QueryRow(context.Background(), "SELECT login FROM admins WHERE login = $1", admin.Login).Scan(&login)
	if login != "" {
		return errors.New("this admin already exists")
	}

	return nil
}
