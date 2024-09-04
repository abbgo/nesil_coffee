package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMails(c *gin.Context) {
	var requestQuery helpers.StandartQuery
	var count uint
	var mails []models.ForMail

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
	if err = db.QueryRow(context.Background(), `SELECT COUNT(id) FROM mails`).Scan(&count); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	rowsMail, err := db.Query(context.Background(), `SELECT id,full_name,email,letter FROM mails 
	ORDER BY created_at DESC LIMIT $1 OFFSET $2`, requestQuery.Limit, offset)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsMail.Close()

	for rowsMail.Next() {
		var mail models.ForMail
		if err := rowsMail.Scan(&mail.ID, &mail.FullName, &mail.Email, &mail.Letter); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		mails = append(mails, mail)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"mails":  mails,
		"count":  count,
	})
}

func DeleteMailByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request parametr - den id alynyar
	ID := c.Param("id")
	if err := helpers.ValidateRecordByID("mails", ID, "NULL", db); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// category - nyn suraty pozulandan son category we onun bilen baglanysykly maglumatlar pozulyar
	_, err = db.Exec(context.Background(), "DELETE FROM mails WHERE id = $1", ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully deleted",
	})
}
