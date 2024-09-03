package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAdmin(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request - den gelen maglumatlar alynyar
	var admin models.Admin
	if err := c.BindJSON(&admin); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if err := models.ValidateRegisterAdmin(admin); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// parol hashlenyan
	hashPassword, err := helpers.HashPassword(admin.Password)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// hemme zat yerbe yer bolsa maglumatlar admins tablisa gosulyar
	_, err = db.Exec(context.Background(), "INSERT INTO admins (login,password) VALUES ($1,$2)", admin.Login, hashPassword)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"login":  admin.Login,
	})
}

func LoginAdmin(c *gin.Context) {
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request - den maglumatlar alynyar
	var admin models.Admin
	if err := c.BindJSON(&admin); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	var id, password string
	db.QueryRow(context.Background(), "SELECT id,password FROM admins WHERE login = $1", admin.Login).Scan(&id, &password)
	// eger request - den gelen telefon belgili admin database - de yok bolsa onda error response edilyar
	if id == "" {
		helpers.HandleError(c, 404, "admin not found")
		return
	}

	// eger admin bar bolsa onda paroly dogry yazypdyrmy yazmandyrmy sol barlanyar
	credentialError := helpers.CheckPassword(admin.Password, password)
	if credentialError != nil {
		helpers.HandleError(c, 400, "invalid credentials")
		return
	}

	// maglumatlar dogry bolsa auth ucin access_toke bilen resfresh_token generate edilyar
	accessTokenString, err := helpers.GenerateAccessToken(admin.Login, id)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": accessTokenString,
		"login":        admin.Login,
		"status":       true,
	})
}
