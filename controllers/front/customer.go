package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterCustomer(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request - den gelen maglumatlar alynyar
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if err := models.ValidateRegisterCustomer(customer); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// parol hashlenyan
	hashPassword, err := helpers.HashPassword(customer.Password)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// hemme zat yerbe yer bolsa maglumatlar admins tablisa gosulyar
	var customerID string
	if err := db.QueryRow(context.Background(), "INSERT INTO customers (login,mail,password) VALUES ($1,$2,$3) RETURNING id",
		customer.Login, customer.Mail, hashPassword,
	).Scan(&customerID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// maglumatlar dogry bolsa auth ucin access_toke bilen resfresh_token generate edilyar
	accessTokenString, err := helpers.GenerateAccessToken(customer.Login, customerID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	customer.Password = ""

	c.JSON(http.StatusOK, gin.H{
		"status":       true,
		"customer":     customer.Login,
		"access_token": accessTokenString,
	})
}
