package models

type Diplom struct {
	ID    string `json:"id"`
	Image string `json:"image" binding:"required"`
}
