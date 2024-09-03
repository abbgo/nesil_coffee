package models

type Gallery struct {
	ID       string `json:"id,omitempty"`
	Media    string `json:"media" binding:"required"`
	MdiaType string `json:"media_type" binding:"required"`
}
