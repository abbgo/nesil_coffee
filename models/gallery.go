package models

import (
	"errors"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
)

type Gallery struct {
	ID        string `json:"id,omitempty"`
	Media     string `json:"media" binding:"required"`
	MediaType string `json:"media_type" binding:"required"`
}

func ValidateCreateGallery(gallery Gallery) error {
	hasMediaType := false
	for _, mediaType := range helpers.MediaTypes {
		if gallery.MediaType == mediaType {
			hasMediaType = true
		}
	}
	if !hasMediaType {
		return errors.New("invalid media type")
	}

	return nil
}

func ValidateUpdateGallery(gallery Gallery) error {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		return err
	}
	defer db.Close()

	if gallery.ID == "" {
		return errors.New("gallery id is required")
	}

	if err := helpers.ValidateRecordByID("galleries", gallery.ID, "NULL", db); err != nil {
		return err
	}

	hasMediaType := false
	for _, mediaType := range helpers.MediaTypes {
		if gallery.MediaType == mediaType {
			hasMediaType = true
		}
	}
	if !hasMediaType {
		return errors.New("invalid media type")
	}

	return nil
}
