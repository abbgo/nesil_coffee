package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateSlider(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var slider models.Slider
	if err := c.BindJSON(&slider); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	hasMediaType := false
	for _, mediaType := range helpers.MediaTypes {
		if slider.MediaType == mediaType {
			hasMediaType = true
		}
	}
	if !hasMediaType {
		helpers.HandleError(c, 400, "invalid media type")
		return
	}

	// eger maglumatlar dogry bolsa db maglumatlar gosulyar
	_, err = db.Exec(context.Background(),
		`INSERT INTO sliders 
		(image_url,title_tm,title_ru,title_en,sub_title_tm,sub_title_ru,sub_title_en,image_hash,media_type) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		slider.Image.URL, slider.TitleTM, slider.TitleRU, slider.TitleEN,
		slider.SubTitleTM, slider.SubTitleRU, slider.SubTitleEN, slider.Image.HashBlur, slider.MediaType,
	)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// maglumat gpsulandan sonra helper images tablisadan media pozulyar
	if err := DeleteImageFromDB(slider.Image.URL); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateSliderByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var slider models.Slider
	if err := c.BindJSON(&slider); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if slider.ID == "" {
		helpers.HandleError(c, 400, "slider id is required")
		return
	}
	if err := helpers.ValidateRecordByID("sliders", slider.ID, "NULL", db); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	hasMediaType := false
	for _, mediaType := range helpers.MediaTypes {
		if slider.MediaType == mediaType {
			hasMediaType = true
		}
	}
	if !hasMediaType {
		helpers.HandleError(c, 400, "invalid media type")
		return
	}

	// database - daki maglumatlary request body - dan gelen maglumatlar bilen calysyas
	_, err = db.Exec(context.Background(),
		`UPDATE sliders SET title_tm=$1 , title_ru=$2 , title_en=$3 , image_url=$4 , sub_title_tm=$5 , sub_title_ru=$6 , 
		sub_title_en=$7 , image_hash=$8 , media_type=$9 WHERE id=$10`,
		slider.TitleTM, slider.TitleRU, slider.TitleEN, slider.Image.URL, slider.SubTitleTM, slider.SubTitleRU, slider.SubTitleEN,
		slider.Image.HashBlur, slider.MediaType, slider.ID,
	)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if err := DeleteImageFromDB(slider.Image.URL); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetSliderByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request parametrden category id alynyar
	sliderID := c.Param("id")

	// database - den request parametr - den gelen id boyunca maglumat cekilyar
	var slider models.Slider
	if err := db.QueryRow(context.Background(),
		`SELECT id,title_tm,title_ru,title_en,sub_title_tm,sub_title_ru,sub_title_en,image_url,image_hash,media_type FROM sliders WHERE id = $1`, sliderID).
		Scan(&slider.ID, &slider.TitleTM, &slider.TitleRU, &slider.TitleEN,
			&slider.SubTitleTM, &slider.SubTitleRU, &slider.SubTitleEN,
			&slider.Image.URL, &slider.Image.HashBlur, &slider.MediaType,
		); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger databse sol maglumat yok bolsa error return edilyar
	if slider.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"slider": slider,
	})
}

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
		`SELECT id,title_tm,title_ru,title_en,sub_title_tm,sub_title_ru,sub_title_en,image_url,image_hash,media_type 
		FROM sliders ORDER BY created_at DESC LIMIT $1 OFFSET $2`, requestQuery.Limit, offset)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsRecipe.Close()

	for rowsRecipe.Next() {
		var slider models.Slider
		if err := rowsRecipe.Scan(&slider.ID, &slider.TitleTM, &slider.TitleRU, &slider.TitleEN,
			&slider.SubTitleTM, &slider.SubTitleRU, &slider.SubTitleEN, &slider.Image.URL,
			&slider.Image.HashBlur, &slider.MediaType,
		); err != nil {
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

func DeleteSliderByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request parametr - den id alynyar
	ID := c.Param("id")
	var slider models.Slider
	db.QueryRow(context.Background(), "SELECT id,image_url FROM sliders WHERE id=$1", ID).Scan(&slider.ID, &slider.Image.URL)
	if slider.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	// local path - dan surat pozulyar
	if err := os.Remove(helpers.ServerPath + slider.Image.URL); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// category - nyn suraty pozulandan son category we onun bilen baglanysykly maglumatlar pozulyar
	_, err = db.Exec(context.Background(), "DELETE FROM sliders WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}
