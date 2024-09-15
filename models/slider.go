package models

type Slider struct {
	ID         string    `json:"id,omitempty"`
	Image      BlurImage `json:"image" binding:"required"`
	TitleTM    string    `json:"title_tm" binding:"required"`
	TitleRU    string    `json:"title_ru" binding:"required"`
	TitleEN    string    `json:"title_en" binding:"required"`
	SubTitleTM string    `json:"sub_title_tm,omitempty"`
	SubTitleRU string    `json:"sub_title_ru,omitempty"`
	SubTitleEN string    `json:"sub_title_en,omitempty"`
}
