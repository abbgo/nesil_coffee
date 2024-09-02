package models

type ProductComposition struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name" binding:"required"`
	Percentage int8   `json:"percentage" binding:"required"`
}
