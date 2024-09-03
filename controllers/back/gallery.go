package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateGallery(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var gallery models.Gallery
	if err := c.BindJSON(&gallery); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// maglumatlar barlanyar
	if err := models.ValidateCreateGallery(gallery); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger maglumatlar dogry bolsa db gosulyar
	_, err = db.Exec(context.Background(), "INSERT INTO galleries (media,media_type) VALUES ($1,$2)",
		gallery.Media, gallery.MediaType)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// maglumat gpsulandan sonra helper images tablisadan media pozulyar
	if err := DeleteImageFromDB(gallery.Media); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateGalleryByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var gallery models.Gallery
	if err := c.BindJSON(&gallery); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// maglumatlar barlanyar
	if err := models.ValidateUpdateGallery(gallery); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE galleries SET media = $1 , media_type = $2 WHERE id = $3",
		gallery.Media, gallery.MediaType, gallery.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if err := DeleteImageFromDB(gallery.Media); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}
