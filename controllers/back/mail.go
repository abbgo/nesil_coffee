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

	rowsMail, err := db.Query(context.Background(),
		`SELECT id,full_name,email,letter,product_id FROM mails 
	ORDER BY created_at DESC LIMIT $1 OFFSET $2`, requestQuery.Limit, offset)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsMail.Close()

	for rowsMail.Next() {
		var mail models.ForMail
		if err := rowsMail.Scan(&mail.ID, &mail.FullName, &mail.Email, &mail.Letter, &mail.ProductID); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		// eger maile degisli haryt bar bolsa sol harydyn maglumatlary alynyar
		if mail.ProductID.String != "" {
			if err := db.QueryRow(context.Background(),
				"SELECT id,name_tm,name_ru,name_en,description_tm,description_ru,description_en,category_id FROM products WHERE id = $1", mail.ProductID).
				Scan(&mail.Product.ID, &mail.Product.NameTM, &mail.Product.NameRU, &mail.Product.NameEN,
					&mail.Product.DescriptionTM, &mail.Product.DescriptionRU, &mail.Product.DescriptionEN,
					&mail.Product.CategoryID); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			// harydyn suraty db - den alynyar
			rowsImage, err := db.Query(context.Background(), "SELECT image FROM product_images WHERE product_id=$1", mail.ProductID)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			defer rowsImage.Close()

			for rowsImage.Next() {
				var image string
				if err := rowsImage.Scan(&image); err != nil {
					helpers.HandleError(c, 400, err.Error())
					return
				}
				mail.Product.Images = append(mail.Product.Images, image)
			}
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
