package models

type Product struct {
	ID          string   `json:"id,omitempty"`
	Name        string   `json:"name" binding:"required"`
	Description string   `json:"description" binding:"required"`
	Images      []string `json:"images" binding:"required"`
	Slug        string   `json:"slug,omitempty"`
}
