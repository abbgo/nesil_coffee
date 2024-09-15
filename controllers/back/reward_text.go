package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRewardText(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var textSlider models.TextSlider
	if err := c.BindJSON(&textSlider); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(
		context.Background(), `INSERT INTO reward_text (description_tm,description_ru,description_en) VALUES ($1,$2,$3)`,
		textSlider.DescriptionTM, textSlider.DescriptionRU, textSlider.DescriptionEN,
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

func UpdateRewardTextByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var textSlider models.TextSlider
	if err := c.BindJSON(&textSlider); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if textSlider.ID == "" {
		helpers.HandleError(c, 400, "id is required")
		return
	}
	var id string
	db.QueryRow(context.Background(), `SELECT id FROM reward_text WHERE id=$1`, textSlider.ID).Scan(&id)
	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(
		context.Background(), `UPDATE reward_text SET description_tm=$1 , description_ru=$2 , description_en=$3 WHERE id=$4`,
		textSlider.DescriptionTM, textSlider.DescriptionRU, textSlider.DescriptionEN, textSlider.ID,
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

func GetOneRewardText(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var textSlider models.TextSlider
	db.QueryRow(context.Background(), `SELECT id,description_tm,description_ru,description_en FROM reward_text LIMIT 1`).
		Scan(&textSlider.ID, &textSlider.DescriptionTM, &textSlider.DescriptionRU, &textSlider.DescriptionEN)
	if textSlider.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":      true,
		"reward_text": textSlider,
	})
}

func DeleteRewardText(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	ID := c.Param("id")
	var id string
	db.QueryRow(context.Background(), `SELECT id FROM reward_text WHERE id=$1`, ID).Scan(&id)
	if id == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	_, err = db.Exec(context.Background(), `DELETE FROM reward_text WHERE id=$1`, ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}
