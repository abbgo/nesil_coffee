package models

type FAQ struct {
	ID            string `json:"id,omitempty"`
	TitleTM       string `json:"title_tm" binding:"required"`
	TitleRU       string `json:"title_ru" binding:"required"`
	TitleEN       string `json:"title_en" binding:"required"`
	DescriptionTM string `json:"description_tm,omitempty"`
	DescriptionRU string `json:"description_ru,omitempty"`
	DescriptionEN string `json:"description_en,omitempty"`
	SlugTM        string `json:"slug_tm,omitempty"`
	SlugRU        string `json:"slug_ru,omitempty"`
	SlugEN        string `json:"slug_en,omitempty"`
}
