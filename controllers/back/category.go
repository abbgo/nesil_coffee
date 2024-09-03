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

func CreateCategory(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var category models.Category
	if err := c.BindJSON(&category); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger maglumatlar dogry bolsa onda categories tablisa maglumatlar gosulyar
	_, err = db.Exec(context.Background(),
		"INSERT INTO categories (name_tm,name_ru,name_en,slug_tm,slug_ru,slug_en) VALUES ($1,$2,$3,$4,$5,$6)",
		category.NameTM, category.NameRU, category.NameEN,
		slug.MakeLang(category.NameTM, "en"), slug.MakeLang(category.NameRU, "en"), slug.MakeLang(category.NameEN, "en"),
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

// func UpdateCategoryByID(c *gin.Context) {
// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	// request body - dan gelen maglumatlar alynyar
// 	var category models.Category
// 	if err := c.BindJSON(&category); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// bind edilen maglumatlar barlanyar
// 	if category.ID == "" {
// 		helpers.HandleError(c, 400, "category id is required")
// 		return
// 	}
// 	if err := helpers.ValidateRecordByID("categories", category.ID, "NULL", db); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// database - daki maglumatlary request body - dan gelen maglumatlar bilen calysyas
// 	_, err = db.Exec(context.Background(),
// 		"UPDATE categories SET name=$1 , image=$2 , description=$3 , slug=$4 WHERE id=$5",
// 		category.Name, category.Image, category.Description, slug.MakeLang(category.Name, "en"), category.ID,
// 	)
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// category - nyn maglumatlary uytgedilenden sonra suraty helper_images tablisa category ucin gosulan surat pozulyar
// 	if category.Image != "" {
// 		if err := DeleteImageFromDB(category.Image); err != nil {
// 			helpers.HandleError(c, 400, err.Error())
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "data successfully updated",
// 	})
// }

// func GetCategoryByID(c *gin.Context) {
// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	// request parametrden category id alynyar
// 	categoryID := c.Param("id")

// 	// database - den request parametr - den gelen id boyunca maglumat cekilyar
// 	var category models.Category
// 	if err := db.QueryRow(context.Background(),
// 		"SELECT id,name,image,description FROM categories WHERE id = $1", categoryID).
// 		Scan(&category.ID, &category.Name, &category.Image, &category.Description); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// eger databse sol maglumat yok bolsa error return edilyar
// 	if category.ID == "" {
// 		helpers.HandleError(c, 404, "record not found")
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":   true,
// 		"category": category,
// 	})
// }

// func GetCategories(c *gin.Context) {
// 	var requestQuery serializations.CategoryQuery
// 	var count uint
// 	var categories []models.Category
// 	deletedAt := `IS NULL`
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

// 	if requestQuery.IsDeleted {
// 		deletedAt = `IS NOT NULL`
// 	}

// 	if requestQuery.Search != "" {
// 		incomingsSarch := slug.MakeLang(c.Query("search"), "en")
// 		search = strings.ReplaceAll(incomingsSarch, "-", " | ")
// 		searchStr = fmt.Sprintf("%%%s%%", search)
// 	}

// 	if requestQuery.Search != "" {
// 		searchQuery = fmt.Sprintf(` AND (to_tsvector(slug) @@ to_tsquery('%s') OR slug LIKE '%s') `, search, searchStr)
// 	}

// 	// database - den maglumatlaryn sany alynyar
// 	queryCount := fmt.Sprintf(`SELECT COUNT(id) FROM categories WHERE deleted_at %s %s`, deletedAt, searchQuery)
// 	if err = db.QueryRow(context.Background(), queryCount).Scan(&count); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// database maglumatlar alynyar
// 	queryCategories := fmt.Sprintf(
// 		`SELECT id,name,image,description FROM categories WHERE deleted_at %s %s ORDER BY created_at DESC LIMIT %d OFFSET %d`,
// 		deletedAt, searchQuery, requestQuery.Limit, offset)
// 	rowsCategory, err := db.Query(context.Background(), queryCategories)
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer rowsCategory.Close()

// 	for rowsCategory.Next() {
// 		var category models.Category
// 		if err := rowsCategory.Scan(&category.ID, &category.Name, &category.Image, &category.Description); err != nil {
// 			helpers.HandleError(c, 400, err.Error())
// 			return
// 		}
// 		categories = append(categories, category)
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":     true,
// 		"categories": categories,
// 		"count":      count,
// 	})
// }

// func DeleteCategoryByID(c *gin.Context) {
// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	// request parametr - den id alynyar
// 	ID := c.Param("id")

// 	// database - de gelen id degisli maglumat barmy sol barlanyar
// 	var id string
// 	var image string
// 	db.QueryRow(context.Background(), "SELECT id,image FROM categories WHERE id = $1", ID).Scan(&id, &image)

// 	// eger database - de gelen id degisli category yok bolsa error return edilyar
// 	if id == "" {
// 		helpers.HandleError(c, 404, "record not found")
// 		return
// 	}

// 	// local path - dan surat pozulyar
// 	if err := os.Remove(helpers.ServerPath + image); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// category - nyn suraty pozulandan son category we onun bilen baglanysykly maglumatlar pozulyar
// 	_, err = db.Exec(context.Background(), "DELETE FROM categories WHERE id = $1", ID)
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "data successfully deleted",
// 	})
// }
