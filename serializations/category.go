package serializations

import "nesil_coffe/helpers"

type CategoryQuery struct {
	helpers.StandartQuery
	Search string `form:"search"`
}
