package models

type ProductComposition struct {
	ID         string `json:"id,omitempty"`
	NameTM     string `json:"name_tm" binding:"required"`
	NameRU     string `json:"name_ru" binding:"required"`
	NameEN     string `json:"name_en" binding:"required"`
	Percentage int8   `json:"percentage" binding:"required"`
}
