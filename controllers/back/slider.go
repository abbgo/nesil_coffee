package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSlider(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var slider models.Slider
	if err := c.BindJSON(&slider); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger maglumatlar dogry bolsa db maglumatlar gosulyar
	_, err = db.Exec(context.Background(),
		`INSERT INTO sliders (image,title_tm,title_ru,title_en,sub_title_tm,sub_title_ru,sub_title_en,description_tm,description_ru,description_en) 
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)`,
		slider.Image, slider.TitleTM, slider.TitleRU, slider.TitleEN,
		slider.SubTitleTM, slider.SubTitleRU, slider.SubTitleEN,
		slider.DescriptionTM, slider.DescriptionRU, slider.DescriptionEN,
	)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// maglumat gpsulandan sonra helper images tablisadan media pozulyar
	if err := DeleteImageFromDB(slider.Image); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

// func UpdateRecipeByID(c *gin.Context) {
// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	// request body - dan gelen maglumatlar alynyar
// 	var recipe models.Recipe
// 	if err := c.BindJSON(&recipe); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	if recipe.ID == "" {
// 		helpers.HandleError(c, 400, "recipe id is required")
// 		return
// 	}
// 	if err := helpers.ValidateRecordByID("recipes", recipe.ID, "NULL", db); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// database - daki maglumatlary request body - dan gelen maglumatlar bilen calysyas
// 	_, err = db.Exec(context.Background(),
// 		`UPDATE recipes SET name_tm=$1 , name_ru=$2 , name_en=$3 , description_tm=$4 , description_ru=$5 , description_en=$6 ,
// 		image=$7 , slug_tm=$8 , slug_ru=$9 , slug_en=$10 WHERE id=$11`,
// 		recipe.NameTM, recipe.NameRU, recipe.NameEN, recipe.DescriptionTM, recipe.DescriptionRU, recipe.DescriptionEN,
// 		recipe.Image, slug.MakeLang(recipe.NameTM, "en"), slug.MakeLang(recipe.NameRU, "en"), slug.MakeLang(recipe.NameEN, "en"),
// 		recipe.ID,
// 	)
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	_, err = db.Exec(context.Background(), "DELETE FROM product_compositions WHERE recipe_id=$1", recipe.ID)
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// eger harydyn duzumi girizilen bolsa ol hem db gosulyar
// 	if len(recipe.Compositions) != 0 {
// 		for _, composition := range recipe.Compositions {
// 			_, err = db.Exec(context.Background(),
// 				"INSERT INTO product_compositions (name_tm,name_ru,name_en,percentage,recipe_id) VALUES ($1,$2,$3,$4,$5)",
// 				composition.NameTM, composition.NameRU, composition.NameEN,
// 				composition.Percentage, recipe.ID,
// 			)
// 			if err != nil {
// 				helpers.HandleError(c, 400, err.Error())
// 				return
// 			}
// 		}
// 	}

// 	if err := DeleteImageFromDB(recipe.Image); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "data successfully updated",
// 	})
// }

// func GetRecipeByID(c *gin.Context) {
// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	// request parametrden category id alynyar
// 	recipeID := c.Param("id")

// 	// database - den request parametr - den gelen id boyunca maglumat cekilyar
// 	var recipe models.Recipe
// 	if err := db.QueryRow(context.Background(),
// 		"SELECT id,name_tm,name_ru,name_en,description_tm,description_ru,description_en,image FROM recipes WHERE id = $1", recipeID).
// 		Scan(&recipe.ID, &recipe.NameTM, &recipe.NameRU, &recipe.NameEN,
// 			&recipe.DescriptionTM, &recipe.DescriptionRU, &recipe.DescriptionEN, &recipe.Image); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// eger databse sol maglumat yok bolsa error return edilyar
// 	if recipe.ID == "" {
// 		helpers.HandleError(c, 404, "record not found")
// 		return
// 	}

// 	// resept duzumi alynyar
// 	rowsComposition, err := db.Query(context.Background(),
// 		"SELECT id,name_tm,name_ru,name_en,percentage FROM product_compositions WHERE recipe_id=$1", recipe.ID,
// 	)
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer rowsComposition.Close()

// 	for rowsComposition.Next() {
// 		var composition models.ProductComposition
// 		if err := rowsComposition.Scan(&composition.ID, &composition.NameTM, &composition.NameRU, &composition.NameEN,
// 			&composition.Percentage); err != nil {
// 			helpers.HandleError(c, 400, err.Error())
// 			return
// 		}

// 		if composition.ID != "" {
// 			recipe.Compositions = append(recipe.Compositions, composition)
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status": true,
// 		"recipe": recipe,
// 	})
// }

// func GetRecipes(c *gin.Context) {
// 	var requestQuery serializations.CategoryQuery
// 	var count uint
// 	var recipes []models.Recipe
// 	var searchQuery, search, searchStr string

// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	// request query - den maglumatlar bind edilyar
// 	if err := c.Bind(&requestQuery); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	// request query - den maglumatlar validate edilyar
// 	if err := helpers.ValidateStructData(&requestQuery); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// limit we page boyunca offset hasaplanyar
// 	offset := requestQuery.Limit * (requestQuery.Page - 1)

// 	if requestQuery.Search != "" {
// 		incomingsSarch := slug.MakeLang(c.Query("search"), "en")
// 		search = strings.ReplaceAll(incomingsSarch, "-", " | ")
// 		searchStr = fmt.Sprintf("%%%s%%", search)
// 	}

// 	if requestQuery.Search != "" {
// 		searchQuery = fmt.Sprintf(` WHERE (to_tsvector(slug_%s) @@ to_tsquery('%s') OR slug_%s LIKE '%s') `, requestQuery.Lang, search, requestQuery.Lang, searchStr)
// 	}

// 	// database - den maglumatlaryn sany alynyar
// 	queryCount := fmt.Sprintf(`SELECT COUNT(id) FROM recipes %s`, searchQuery)
// 	if err = db.QueryRow(context.Background(), queryCount).Scan(&count); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// database maglumatlar alynyar
// 	queryRecipes := fmt.Sprintf(
// 		`SELECT id,name_tm,name_ru,name_en,description_tm,description_ru,description_en,image FROM recipes
// 			%s ORDER BY created_at DESC LIMIT %d OFFSET %d`,
// 		searchQuery, requestQuery.Limit, offset)
// 	rowsRecipe, err := db.Query(context.Background(), queryRecipes)
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer rowsRecipe.Close()

// 	for rowsRecipe.Next() {
// 		var recipe models.Recipe
// 		if err := rowsRecipe.Scan(&recipe.ID, &recipe.NameTM, &recipe.NameRU,
// 			&recipe.NameEN, &recipe.DescriptionTM, &recipe.DescriptionRU, &recipe.DescriptionEN, &recipe.Image); err != nil {
// 			helpers.HandleError(c, 400, err.Error())
// 			return
// 		}

// 		// harydyn duzumi alynyar
// 		rowsComposition, err := db.Query(context.Background(),
// 			"SELECT id,name_tm,name_ru,name_en,percentage FROM product_compositions WHERE recipe_id=$1", recipe.ID,
// 		)
// 		if err != nil {
// 			helpers.HandleError(c, 400, err.Error())
// 			return
// 		}
// 		defer rowsComposition.Close()

// 		for rowsComposition.Next() {
// 			var composition models.ProductComposition
// 			if err := rowsComposition.Scan(&composition.ID,
// 				&composition.NameTM, &composition.NameRU, &composition.NameEN, &composition.Percentage); err != nil {
// 				helpers.HandleError(c, 400, err.Error())
// 				return
// 			}

// 			if composition.ID != "" {
// 				recipe.Compositions = append(recipe.Compositions, composition)
// 			}
// 		}

// 		recipes = append(recipes, recipe)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"recipes": recipes,
// 		"count":   count,
// 	})
// }

// func DeleteRecipeByID(c *gin.Context) {
// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	// request parametr - den id alynyar
// 	ID := c.Param("id")
// 	var recipe models.Recipe
// 	db.QueryRow(context.Background(), "SELECT id,image FROM recipes WHERE id=$1", ID).Scan(&recipe.ID, &recipe.Image)
// 	if recipe.ID == "" {
// 		helpers.HandleError(c, 404, "record not found")
// 		return
// 	}

// 	// local path - dan surat pozulyar
// 	if err := os.Remove(helpers.ServerPath + recipe.Image); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// category - nyn suraty pozulandan son category we onun bilen baglanysykly maglumatlar pozulyar
// 	_, err = db.Exec(context.Background(), "DELETE FROM recipes WHERE id = $1", ID)
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "data successfully deleted",
// 	})
// }
