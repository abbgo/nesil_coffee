package models

type RewardText struct {
	ID            string `json:"id"`
	DescriptionTM string `json:"description_tm" binding:"required"`
	DescriptionRU string `json:"description_ru" binding:"required"`
	DescriptionEN string `json:"description_en" binding:"required"`
}
