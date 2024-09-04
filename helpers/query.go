package helpers

type StandartQuery struct {
	Limit int `form:"limit" validate:"required,min=10,max=100"`
	Page  int `form:"page" validate:"required,min=1"`
}
