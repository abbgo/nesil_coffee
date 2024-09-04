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

func GetGalleryByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request parametrden category id alynyar
	galleryID := c.Param("id")

	// database - den request parametr - den gelen id boyunca maglumat cekilyar
	var gallery models.Gallery
	if err := db.QueryRow(context.Background(),
		"SELECT id,media,media_type FROM galleries WHERE id = $1", galleryID).
		Scan(&gallery.ID, &gallery.Media, &gallery.MediaType); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger databse sol maglumat yok bolsa error return edilyar
	if gallery.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"gallery": gallery,
	})
}

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

func DeleteGalleryByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request parametr - den id alynyar
	ID := c.Param("id")
	var gallery models.Gallery
	db.QueryRow(context.Background(), "SELECT id,media FROM galleries WHERE id=$1", ID).Scan(&gallery.ID, &gallery.Media)
	if gallery.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	// local path - dan surat pozulyar
	if err := os.Remove(helpers.ServerPath + gallery.Media); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// category - nyn suraty pozulandan son category we onun bilen baglanysykly maglumatlar pozulyar
	_, err = db.Exec(context.Background(), "DELETE FROM galleries WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}
