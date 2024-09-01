package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func CreateCategory(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger maglumatlar dogry bolsa onda categories tablisa maglumatlar gosulyar we gosulandan son gosulan maglumatyn id - si return edilyar
	_, err = db.Exec(context.Background(),
		"INSERT INTO categories (name,image,description,slug) VALUES ($1,$2,$3,$4)",
		category.Name, category.Image, category.Description, slug.MakeLang(category.Name, "en"),
	)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// category - nyn maglumatlary gosulandan sonra suraty helper_images tablisa category ucin gosulan surat pozulyar
	if category.Image != "" {
		if err := DeleteImageFromDB(category.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateCategoryByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// bind edilen maglumatlar barlanyar
	if category.ID == "" {
		helpers.HandleError(c, 400, "category id is required")
		return
	}
	if err := helpers.ValidateRecordByID("categories", category.ID, "NULL", db); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database - daki maglumatlary request body - dan gelen maglumatlar bilen calysyas
	_, err = db.Exec(context.Background(),
		"UPDATE categories SET name=$1 , image=$2 , description=$3 , slug=$4 WHERE id=$5",
		category.Name, category.Image, category.Description, slug.MakeLang(category.Name, "en"), category.ID,
	)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}
