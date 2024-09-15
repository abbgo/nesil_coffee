package models

type BlurImage struct {
	URL      string `json:"url" binding:"required"`
	HashBlur string `json:"hash_blur" binding:"required"`
}
