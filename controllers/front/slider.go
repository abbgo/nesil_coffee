package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSliders(c *gin.Context) {
	var requestQuery helpers.StandartQuery
	var count uint
	var sliders []models.Slider

	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request query - den maglumatlar bind edilyar
	if err := c.Bind(&requestQuery); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	// request query - den maglumatlar validate edilyar
	if err := helpers.ValidateStructData(&requestQuery); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// limit we page boyunca offset hasaplanyar
	offset := requestQuery.Limit * (requestQuery.Page - 1)

	// database - den maglumatlaryn sany alynyar
	if err = db.QueryRow(context.Background(), `SELECT COUNT(id) FROM sliders`).Scan(&count); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database maglumatlar alynyar

	rowsRecipe, err := db.Query(context.Background(),
		`SELECT id,title_tm,title_ru,title_en,sub_title_tm,sub_title_ru,sub_title_en,description_tm,description_ru,description_en,
		image_url,image_hash FROM sliders ORDER BY created_at DESC LIMIT $1 OFFSET $2`, requestQuery.Limit, offset)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsRecipe.Close()

	for rowsRecipe.Next() {
		var slider models.Slider
		if err := rowsRecipe.Scan(&slider.ID, &slider.TitleTM, &slider.TitleRU, &slider.TitleEN,
			&slider.SubTitleTM, &slider.SubTitleRU, &slider.SubTitleEN,
			&slider.DescriptionTM, &slider.DescriptionRU, &slider.DescriptionEN, &slider.Image.URL, &slider.Image.HashBlur); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		sliders = append(sliders, slider)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"sliders": sliders,
		"count":   count,
	})
}
