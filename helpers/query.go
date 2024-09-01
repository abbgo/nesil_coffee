package helpers

type StandartQuery struct {
	IsDeleted bool `form:"is_deleted"`
	Limit     int  `form:"limit" validate:"required,min=10,max=100"`
	Page      int  `form:"page" validate:"required,min=1"`
}
