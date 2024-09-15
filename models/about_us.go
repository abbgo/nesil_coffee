package models

type AboutUs struct {
	ID            string `json:"id"`
	TitleTM       string `json:"title_tm" binding:"required"`
	TitleRU       string `json:"title_ru" binding:"required"`
	TitleEN       string `json:"title_en" binding:"required"`
	DescriptionTM string `json:"description_tm" binding:"required"`
	DescriptionRU string `json:"description_ru" binding:"required"`
	DescriptionEN string `json:"description_en" binding:"required"`
}
