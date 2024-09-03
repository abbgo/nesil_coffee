package models

type Category struct {
	ID     string `json:"id,omitempty"`
	NameTM string `json:"name_tm" binding:"required"`
	NameRU string `json:"name_ru" binding:"required"`
	NameEN string `json:"name_en" binding:"required"`
	SlugTM string `json:"slug_tm,omitempty"`
	SlugRU string `json:"slug_ru,omitempty"`
	SlugEN string `json:"slug_en,omitempty"`
}
