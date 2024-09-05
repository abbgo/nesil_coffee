package models

import (
	"nesil_coffe/serializations"

	"gopkg.in/guregu/null.v4"
)

type ForMail struct {
	ID        string                        `json:"id,omitempty"`
	FullName  string                        `json:"full_name" binding:"required,min=3"`
	Email     string                        `json:"email" binding:"email"`
	Letter    string                        `json:"letter" binding:"required,min=3"`
	ProductID null.String                   `json:"product_id,omitempty"`
	Product   serializations.ProductForMail `json:"product,omitempty"`
}
