package models

import (
	"errors"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
)

type Product struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Images      []string `json:"images" binding:"required"`
	CategoryID  string   `json:"category_id" binding:"required"`
	Slug        string   `json:"slug,omitempty"`
}

func ValidateCreateProduct(product Product) error {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// harydyn suratlary barlanyar
	if len(product.Images) == 0 {
		return errors.New("images of product is requred")
	}

	// harydyn kategoriyasy barlanyar
	if err := helpers.ValidateRecordByID("categories", product.CategoryID, "NULL", db); err != nil {
		return err
	}

	return nil
}
