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

func GetRecipeByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request parametrden category id alynyar
	recipeID := c.Param("id")

	// database - den request parametr - den gelen id boyunca maglumat cekilyar
	var recipe models.Recipe
	if err := db.QueryRow(context.Background(),
		"SELECT id,name_tm,name_ru,name_en,description_tm,description_ru,description_en,image FROM recipes WHERE id = $1", recipeID).
		Scan(&recipe.ID, &recipe.NameTM, &recipe.NameRU, &recipe.NameEN,
			&recipe.DescriptionTM, &recipe.DescriptionRU, &recipe.DescriptionEN, &recipe.Image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger databse sol maglumat yok bolsa error return edilyar
	if recipe.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	// resept duzumi alynyar
	rowsComposition, err := db.Query(context.Background(),
		"SELECT id,name_tm,name_ru,name_en,percentage FROM product_compositions WHERE recipe_id=$1", recipe.ID,
	)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsComposition.Close()

	for rowsComposition.Next() {
		var composition models.ProductComposition
		if err := rowsComposition.Scan(&composition.ID, &composition.NameTM, &composition.NameRU, &composition.NameEN,
			&composition.Percentage); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		if composition.ID != "" {
			recipe.Compositions = append(recipe.Compositions, composition)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"recipe": recipe,
	})
}

func GetRecipes(c *gin.Context) {
	var requestQuery serializations.CategoryQuery
	var count uint
	var recipes []models.Recipe
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
	queryCount := fmt.Sprintf(`SELECT COUNT(id) FROM recipes %s`, searchQuery)
	if err = db.QueryRow(context.Background(), queryCount).Scan(&count); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database maglumatlar alynyar
	queryRecipes := fmt.Sprintf(
		`SELECT id,name_tm,name_ru,name_en,description_tm,description_ru,description_en,image FROM recipes 
			%s ORDER BY created_at DESC LIMIT %d OFFSET %d`,
		searchQuery, requestQuery.Limit, offset)
	rowsRecipe, err := db.Query(context.Background(), queryRecipes)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsRecipe.Close()

	for rowsRecipe.Next() {
		var recipe models.Recipe
		if err := rowsRecipe.Scan(&recipe.ID, &recipe.NameTM, &recipe.NameRU,
			&recipe.NameEN, &recipe.DescriptionTM, &recipe.DescriptionRU, &recipe.DescriptionEN, &recipe.Image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		// harydyn duzumi alynyar
		rowsComposition, err := db.Query(context.Background(),
			"SELECT id,name_tm,name_ru,name_en,percentage FROM product_compositions WHERE recipe_id=$1", recipe.ID,
		)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowsComposition.Close()

		for rowsComposition.Next() {
			var composition models.ProductComposition
			if err := rowsComposition.Scan(&composition.ID,
				&composition.NameTM, &composition.NameRU, &composition.NameEN, &composition.Percentage); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}

			if composition.ID != "" {
				recipe.Compositions = append(recipe.Compositions, composition)
			}
		}

		recipes = append(recipes, recipe)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"recipes": recipes,
		"count":   count,
	})
}
