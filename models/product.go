package models

import (
	"errors"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
)

type Product struct {
	ID            string               `json:"id,omitempty"`
	NameTM        string               `json:"name_tm" binding:"required"`
	NameRU        string               `json:"name_ru" binding:"required"`
	NameEN        string               `json:"name_en" binding:"required"`
	DescriptionTM string               `json:"description_tm" binding:"required"`
	DescriptionRU string               `json:"description_ru" binding:"required"`
	DescriptionEN string               `json:"description_en" binding:"required"`
	Images        []string             `json:"images" binding:"required"`
	CategoryID    string               `json:"category_id" binding:"required"`
	Compositions  []ProductComposition `json:"compositions"`
	SlugTM        string               `json:"slug_tm,omitempty"`
	SlugRU        string               `json:"slug_ru,omitempty"`
	SlugEN        string               `json:"slug_en,omitempty"`
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

func ValidateUpdateProduct(product Product) error {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	// bind edilen maglumatlar barlanyar
	if product.ID == "" {
		return errors.New("product id is required")
	}
	if err := helpers.ValidateRecordByID("products", product.ID, "NULL", db); err != nil {
		return err
	}

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
