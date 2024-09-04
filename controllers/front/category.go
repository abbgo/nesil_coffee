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

func GetCategories(c *gin.Context) {
	var requestQuery serializations.CategoryQuery
	var count uint
	var categories []models.Category
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
	queryCount := fmt.Sprintf(`SELECT COUNT(id) FROM categories %s`, searchQuery)
	if err = db.QueryRow(context.Background(), queryCount).Scan(&count); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database maglumatlar alynyar
	queryCategories := fmt.Sprintf(
		`SELECT id,name_tm,name_ru,name_en FROM categories %s ORDER BY created_at DESC LIMIT %d OFFSET %d`,
		searchQuery, requestQuery.Limit, offset)
	rowsCategory, err := db.Query(context.Background(), queryCategories)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsCategory.Close()

	for rowsCategory.Next() {
		var category models.Category
		if err := rowsCategory.Scan(&category.ID, &category.NameTM, &category.NameRU, &category.NameEN); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		categories = append(categories, category)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     true,
		"categories": categories,
		"count":      count,
	})
}
