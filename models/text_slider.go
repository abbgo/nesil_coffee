package models

type TextSlider struct {
	ID            string `json:"id"`
	DescriptionTM string `json:"descriptions_tm" binding:"required"`
	DescriptionRU string `json:"descriptions_ru" binding:"required"`
	DescriptionEN string `json:"descriptions_en" binding:"required"`
}
