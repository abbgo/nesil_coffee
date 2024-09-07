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

func UpdateFAQByID(c *gin.Context) {
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

	if faq.ID == "" {
		helpers.HandleError(c, 400, "faq id is required")
		return
	}
	if err := helpers.ValidateRecordByID("faqs", faq.ID, "NULL", db); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database - daki maglumatlary request body - dan gelen maglumatlar bilen calysyas
	_, err = db.Exec(context.Background(),
		`UPDATE faqs SET title_tm=$1 , title_ru=$2 , title_en=$3 , description_tm=$4 , description_ru=$5 , description_en=$6 ,
		slug_tm=$7 , slug_ru=$8 , slug_en=$9 WHERE id=$10`,
		faq.TitleTM, faq.TitleRU, faq.TitleEN, faq.DescriptionTM, faq.DescriptionRU, faq.DescriptionEN,
		slug.MakeLang(faq.TitleTM, "en"), slug.MakeLang(faq.TitleRU, "en"), slug.MakeLang(faq.TitleEN, "en"),
		faq.ID,
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

func GetFAQByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request parametrden category id alynyar
	faqID := c.Param("id")

	// database - den request parametr - den gelen id boyunca maglumat cekilyar
	var faq models.FAQ
	if err := db.QueryRow(context.Background(),
		`SELECT id,title_tm,title_ru,title_en,description_tm,description_ru,description_en
		FROM faqs WHERE id = $1`, faqID).
		Scan(&faq.ID, &faq.TitleTM, &faq.TitleRU, &faq.TitleEN,
			&faq.DescriptionTM, &faq.DescriptionRU, &faq.DescriptionEN,
		); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger databse sol maglumat yok bolsa error return edilyar
	if faq.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"faq":    faq,
	})
}
