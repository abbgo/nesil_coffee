package models

type Recipe struct {
	ID            string               `json:"id,omitempty"`
	NameTM        string               `json:"name_tm" binding:"required"`
	NameRU        string               `json:"name_ru" binding:"required"`
	NameEN        string               `json:"name_en" binding:"required"`
	DescriptionTM string               `json:"description_tm" binding:"required"`
	DescriptionRU string               `json:"description_ru" binding:"required"`
	DescriptionEN string               `json:"description_en" binding:"required"`
	Image         string               `json:"image" binding:"required"`
	Compositions  []ProductComposition `json:"compositions"`
	SlugTM        string               `json:"slug_tm,omitempty"`
	SlugRU        string               `json:"slug_ru,omitempty"`
	SlugEN        string               `json:"slug_en,omitempty"`
}
