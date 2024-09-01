package models

type Category struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name" binding:"required"`
	Image       string `json:"image" binding:"required"`
	Description string `json:"description,omitempty"`
	Slug        string `json:"slug,omitempty"`
}
