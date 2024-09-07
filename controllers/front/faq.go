package controllers

import (
	"context"
	"fmt"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"nesil_coffe/serializations"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

func GetFAQs(c *gin.Context) {
	var requestQuery serializations.CategoryQuery
	var count uint
	var faqs []models.FAQ
	var searchQuery, search, searchStr string

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

	if requestQuery.Search != "" {
		incomingsSarch := slug.MakeLang(c.Query("search"), "en")
		search = strings.ReplaceAll(incomingsSarch, "-", " | ")
		searchStr = fmt.Sprintf("%%%s%%", search)
	}

	if requestQuery.Search != "" {
		searchQuery = fmt.Sprintf(` WHERE (to_tsvector(slug_%s) @@ to_tsquery('%s') OR slug_%s LIKE '%s') `, requestQuery.Lang, search, requestQuery.Lang, searchStr)
	}

	// database - den maglumatlaryn sany alynyar
	queryCount := fmt.Sprintf(`SELECT COUNT(id) FROM faqs %s`, searchQuery)
	if err = db.QueryRow(context.Background(), queryCount).Scan(&count); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database maglumatlar alynyar
	queryFAQs := fmt.Sprintf(`SELECT id,title_tm,title_ru,title_en,description_tm,description_ru,description_en 
	 FROM faqs %s ORDER BY created_at DESC LIMIT %d OFFSET %d`, searchQuery, requestQuery.Limit, offset)
	rowsRecipe, err := db.Query(context.Background(), queryFAQs)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsRecipe.Close()

	for rowsRecipe.Next() {
		var faq models.FAQ
		if err := rowsRecipe.Scan(&faq.ID, &faq.TitleTM, &faq.TitleRU, &faq.TitleEN,
			&faq.DescriptionTM, &faq.DescriptionRU, &faq.DescriptionEN,
		); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		faqs = append(faqs, faq)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"faqs":   faqs,
		"count":  count,
	})
}
