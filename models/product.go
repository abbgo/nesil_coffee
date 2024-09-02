package models

type Product struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Slug        string `json:"slug,omitempty"`
}
