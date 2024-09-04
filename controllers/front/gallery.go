package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetGalleries(c *gin.Context) {
	var requestQuery helpers.StandartQuery
	var count uint
	var galleries []models.Gallery

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
	if err = db.QueryRow(context.Background(), `SELECT COUNT(id) FROM galleries`).Scan(&count); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database maglumatlar alynyar
	rowsGallery, err := db.Query(context.Background(),
		`SELECT id,media,media_type FROM galleries ORDER BY created_at DESC LIMIT $1 OFFSET $2`,
		requestQuery.Limit, offset)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsGallery.Close()

	for rowsGallery.Next() {
		var gallery models.Gallery
		if err := rowsGallery.Scan(&gallery.ID, &gallery.Media, &gallery.MediaType); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		galleries = append(galleries, gallery)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    true,
		"galleries": galleries,
		"count":     count,
	})
}
