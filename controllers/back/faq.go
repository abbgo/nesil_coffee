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

func CreateFAQ(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var faq models.FAQ
	if err := c.BindJSON(&faq); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger maglumatlar dogry bolsa db maglumatlar gosulyar
	_, err = db.Exec(context.Background(),
		`INSERT INTO faqs (title_tm,title_ru,title_en,description_tm,description_ru,description_en,slug_tm,slug_ru,slug_en) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		faq.TitleTM, faq.TitleRU, faq.TitleEN,
		faq.DescriptionTM, faq.DescriptionRU, faq.DescriptionEN,
		slug.MakeLang(faq.TitleTM, "en"), slug.MakeLang(faq.TitleRU, "en"), slug.MakeLang(faq.TitleEN, "en"),
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
