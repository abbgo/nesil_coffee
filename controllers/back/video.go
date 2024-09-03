package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AddOrUpdateVideo(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var path, file_name string

	imageType := c.Query("image_type")

	oldImage := c.PostForm("old_path")
	if oldImage != "" {
		if err := os.Remove(helpers.ServerPath + oldImage); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		_, err := db.Exec(context.Background(), "DELETE FROM helper_images WHERE image = $1", oldImage)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	switch imageType {
	case "product":
		path = "product"
		file_name = "image"
	case "media":
		path = "media"
		file_name = "image"
	default:
		helpers.HandleError(c, 400, "invalid image")
		return
	}

	image, err := helpers.FileUpload(file_name, path, c)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "INSERT INTO helper_images (image) VALUES ($1)", image)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"image":  image,
	})
}
