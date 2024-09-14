package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDiplom(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var diplom models.Diplom
	if err := c.BindJSON(&diplom); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger maglumatlar dogry bolsa db gosulyar
	_, err = db.Exec(context.Background(), "INSERT INTO diploms (image) VALUES ($1)",
		diplom.Image,
	)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// maglumat gpsulandan sonra helper images tablisadan media pozulyar
	if err := DeleteImageFromDB(diplom.Image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateDiplomByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var diplom models.Diplom
	if err := c.BindJSON(&diplom); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "UPDATE diploms SET image = $1 WHERE id = $2",
		diplom.Image, diplom.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if err := DeleteImageFromDB(diplom.Image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetDiplomByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request parametrden category id alynyar
	diplomID := c.Param("id")

	// database - den request parametr - den gelen id boyunca maglumat cekilyar
	var diplom models.Diplom
	if err := db.QueryRow(context.Background(),
		"SELECT id,image FROM diploms WHERE id = $1", diplomID).
		Scan(&diplom.ID, &diplom.Image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger databse sol maglumat yok bolsa error return edilyar
	if diplom.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"diplom": diplom,
	})
}

func GetDiploms(c *gin.Context) {
	var requestQuery helpers.StandartQuery
	var count uint
	var diploms []models.Diplom

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
	if err = db.QueryRow(context.Background(), `SELECT COUNT(id) FROM diploms`).Scan(&count); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database maglumatlar alynyar
	rowsDiplom, err := db.Query(context.Background(),
		`SELECT id,image FROM diploms ORDER BY created_at DESC LIMIT $1 OFFSET $2`,
		requestQuery.Limit, offset)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsDiplom.Close()

	for rowsDiplom.Next() {
		var diplom models.Diplom
		if err := rowsDiplom.Scan(&diplom.ID, &diplom.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		diploms = append(diploms, diplom)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"diploms": diploms,
		"count":   count,
	})
}
