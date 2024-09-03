package models

type Admin struct {
	ID       string `json:"id,omitempty"`
	Login    string `json:"login,omitempty" binding:"required"`
	Password string `json:"password,omitempty"`
}
