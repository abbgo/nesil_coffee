package models

type ForMail struct {
	ID        string `json:"id,omitempty"`
	FullName  string `json:"full_name" binding:"required,min=3"`
	Email     string `json:"email" binding:"email"`
	Letter    string `json:"letter" binding:"required,min=3"`
	ProductID string `json:"product_id,omitempty"`
}
