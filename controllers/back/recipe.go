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

func CreateRecipe(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var recipe models.Recipe
	if err := c.BindJSON(&recipe); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger maglumatlar dogry bolsa db maglumatlar gosulyar we gosulandan son gosulan maglumatyn id - si return edilyar
	var recipeID string
	if err := db.QueryRow(context.Background(),
		`INSERT INTO recipes (name_tm,name_ru,name_en,description_tm,description_ru,description_en,image,slug_tm,slug_ru,slug_en) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`,
		recipe.NameTM, recipe.NameRU, recipe.NameEN,
		recipe.DescriptionTM, recipe.DescriptionRU, recipe.DescriptionEN,
		recipe.Image, slug.MakeLang(recipe.NameTM, "en"), slug.MakeLang(recipe.NameRU, "en"), slug.MakeLang(recipe.NameEN, "en"),
	).Scan(&recipeID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger reseptenin duzumi db gosulyar
	if len(recipe.Compositions) != 0 {
		for _, composition := range recipe.Compositions {
			_, err = db.Exec(context.Background(),
				"INSERT INTO product_compositions (name_tm,name_ru,name_en,percentage,recipe_id) VALUES ($1,$2,$3,$4,$5)",
				composition.NameTM, composition.NameRU, composition.NameEN,
				composition.Percentage, recipeID,
			)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	}

	// maglumat gpsulandan sonra helper images tablisadan media pozulyar
	if err := DeleteImageFromDB(recipe.Image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}
