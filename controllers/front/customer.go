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

func LoginCustomer(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request - den maglumatlar alynyar
	var customer models.LoginCustomer
	if err := c.BindJSON(&customer); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	returnCustomer, err := models.ValidateLoginCustomer(customer)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// maglumatlar dogry bolsa auth ucin access_toke bilen resfresh_token generate edilyar
	accessTokenString, err := helpers.GenerateAccessToken(customer.Login, returnCustomer.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessTokenString,
		"customer":     returnCustomer,
		"status":       true,
	})
}

func UpdateCustomer(c *gin.Context) {
	customerID, hasID := c.Get("id")
	if !hasID {
		helpers.HandleError(c, 400, "customer id is required")
		return
	}

	var ok bool
	customer_id, ok := customerID.(string)
	if !ok {
		helpers.HandleError(c, 400, "customer id must be string")
		return
	}

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - den maglumatlar alynyar
	var customer models.Customer
	if err := c.BindJSON(&customer); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// parol hashlenyan
	hashPassword, err := helpers.HashPassword(customer.Password)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger customer database - de bar bolsa onda onun maglumatlary request body - dan gelen maglumatlar bilen update edilyar
	_, err = db.Exec(context.Background(), "UPDATE customers SET login = $1 , mail = $2 , password = $3 WHERE id = $4",
		customer.Login, hashPassword, customer.Mail, customer_id)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetCustomer(c *gin.Context) {
	customerID, hasID := c.Get("id")
	if !hasID {
		helpers.HandleError(c, 400, "customer id is required")
		return
	}

	var ok bool
	customer_id, ok := customerID.(string)
	if !ok {
		helpers.HandleError(c, 400, "customer id must be string")
		return
	}

	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	var customer models.Customer
	db.QueryRow(context.Background(), "SELECT id,login,mail FROM customers WHERE id = $1", customer_id).
		Scan(&customer.ID, &customer.Login, &customer.Mail)
	// eger request - den gelen telefon belgili admin database - de yok bolsa onda error response edilyar
	if customer.ID == "" {
		helpers.HandleError(c, 404, "customer not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"customer": customer,
	})
}
