package models

import (
	"errors"
	"nesil_coffe/helpers"
)

type Gallery struct {
	ID       string `json:"id,omitempty"`
	Media    string `json:"media" binding:"required"`
	MdiaType string `json:"media_type" binding:"required"`
}

func ValidateCreateGallery(gallery Gallery) error {
	hasMediaType := false
	for _, mediaType := range helpers.MediaTypes {
		if gallery.MdiaType == mediaType {
			hasMediaType = true
		}
	}
	if !hasMediaType {
		return errors.New("invalid media type")
	}

	return nil
}
