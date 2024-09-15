package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAboutUs(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var aboutUs models.AboutUs
	if err := c.BindJSON(&aboutUs); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(
		context.Background(), `INSERT INTO about_us (title_tm,title_ru,title_en,description_tm,description_ru,description_en) 
		VALUES ($1,$2,$3,$4,$5,$6)`,
		aboutUs.TitleTM, aboutUs.TitleRU, aboutUs.TitleEN,
		aboutUs.DescriptionTM, aboutUs.DescriptionRU, aboutUs.DescriptionEN,
	)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateAboutUsByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var aboutUs models.AboutUs
	if err := c.BindJSON(&aboutUs); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if aboutUs.ID == "" {
		helpers.HandleError(c, 400, "id is required")
		return
	}
	var id string
	db.QueryRow(context.Background(), `SELECT id FROM about_us WHERE id=$1`, aboutUs.ID).Scan(&id)
	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(
		context.Background(), `UPDATE about_us SET title_tm=$1 , title_ru=$2 , title_en=$3,
		 description_tm=$4 , description_ru=$5 , description_en=$6 WHERE id=$7`,
		aboutUs.TitleTM, aboutUs.TitleRU, aboutUs.TitleEN,
		aboutUs.DescriptionTM, aboutUs.DescriptionRU, aboutUs.DescriptionEN, aboutUs.ID,
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

func GetOneAboutUs(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var aboutUs models.AboutUs
	db.QueryRow(context.Background(), `SELECT id,title_tm,title_ru,title_en,description_tm,description_ru,description_en FROM 
	about_us LIMIT 1`).
		Scan(&aboutUs.ID, &aboutUs.TitleTM, &aboutUs.TitleRU, &aboutUs.TitleEN,
			&aboutUs.DescriptionTM, &aboutUs.DescriptionRU, &aboutUs.DescriptionEN,
		)
	if aboutUs.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"about_us": aboutUs,
	})
}

func DeleteAboutUs(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	ID := c.Param("id")
	var id string
	db.QueryRow(context.Background(), `SELECT id FROM about_us WHERE id=$1`, ID).Scan(&id)
	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), `DELETE FROM about_us WHERE id=$1`, ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}
