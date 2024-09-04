package models

type ForMail struct {
	FullName string `json:"full_name" binding:"required,min=3"`
	Email    string `json:"email" binding:"email"`
	Letter   string `json:"letter" binding:"required,min=3"`
}
