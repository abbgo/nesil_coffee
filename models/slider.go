package models

type Slider struct {
	ID         string    `json:"id,omitempty"`
	Image      BlurImage `json:"image"`
	TitleTM    string    `json:"title_tm"`
	TitleRU    string    `json:"title_ru"`
	TitleEN    string    `json:"title_en"`
	SubTitleTM string    `json:"sub_title_tm,omitempty"`
	SubTitleRU string    `json:"sub_title_ru,omitempty"`
	SubTitleEN string    `json:"sub_title_en,omitempty"`
	MediaType  string    `json:"media_type"`
}
