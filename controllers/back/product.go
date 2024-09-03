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
	"github.com/lib/pq"
)

func CreateProduct(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if err := models.ValidateCreateProduct(product); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger maglumatlar dogry bolsa db maglumatlar gosulyar we gosulandan son gosulan maglumatyn id - si return edilyar
	var productID string
	if err := db.QueryRow(context.Background(),
		`INSERT INTO products (name_tm,name_ru,name_en,description_tm,description_ru,description_en,category_id,slug_tm,slug_ru,slug_en) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id`,
		product.NameTM, product.NameRU, product.NameEN,
		product.DescriptionTM, product.DescriptionRU, product.DescriptionEN,
		product.CategoryID, slug.MakeLang(product.NameTM, "en"), slug.MakeLang(product.NameRU, "en"), slug.MakeLang(product.NameEN, "en"),
	).Scan(&productID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// haryt db gosulandan son harydyn suratlary db gosulyar
	_, err = db.Exec(context.Background(), "INSERT INTO product_images (image,product_id) VALUES (unnest($1::VARCHAR[]),$2)",
		pq.Array(product.Images), productID,
	)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger harydyn duzumi girizilen bolsa ol hem db gosulyar
	if len(product.Compositions) != 0 {
		for _, composition := range product.Compositions {
			_, err = db.Exec(context.Background(),
				"INSERT INTO product_compositions (name_tm,name_ru,name_en,percentage,product_id) VALUES ($1,$2,$3,$4,$5)",
				composition.NameTM, composition.NameRU, composition.NameEN,
				composition.Percentage, productID,
			)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	}

	// harydyn suratlary db gosulandan son helper_images tablisadan suratlar pozulyar
	_, err = db.Exec(context.Background(), "DELETE FROM helper_images WHERE image = ANY($1::VARCHAR[])", product.Images)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully added",
	})
}

func UpdateProductByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request body - dan gelen maglumatlar alynyar
	var product models.Product
	if err := c.BindJSON(&product); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	if err := models.ValidateUpdateProduct(product); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database - daki maglumatlary request body - dan gelen maglumatlar bilen calysyas
	_, err = db.Exec(context.Background(),
		`UPDATE products SET name_tm=$1 , name_ru=$2 , name_en=$3 , description_tm=$4 , description_ru=$5 , description_en=$6 , 
		category_id=$7 , slug_tm=$8 , slug_ru=$9 , slug_en=$10 WHERE id=$11`,
		product.NameTM, product.NameRU, product.NameEN, product.DescriptionTM, product.DescriptionRU, product.DescriptionEN,
		product.CategoryID, slug.MakeLang(product.NameTM, "en"), slug.MakeLang(product.NameRU, "en"), slug.MakeLang(product.NameEN, "en"),
		product.ID,
	)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// harydyn maglumatlary uytgedilenson suratlary uytgedilyar
	_, err = db.Exec(context.Background(), "DELETE FROM product_images WHERE product_id=$1", product.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "INSERT INTO product_images (image,product_id) VALUES (unnest($1::VARCHAR[]),$2)",
		pq.Array(product.Images), product.ID,
	)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	_, err = db.Exec(context.Background(), "DELETE FROM product_compositions WHERE product_id=$1", product.ID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger harydyn duzumi girizilen bolsa ol hem db gosulyar
	if len(product.Compositions) != 0 {
		for _, composition := range product.Compositions {
			_, err = db.Exec(context.Background(),
				"INSERT INTO product_compositions (name_tm,name_ru,name_en,percentage,product_id) VALUES ($1,$2,$3,$4,$5)",
				composition.NameTM, composition.NameRU, composition.NameEN,
				composition.Percentage, product.ID,
			)
			if err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
		}
	}

	// harydyn suratlary db gosulandan son helper_images tablisadan suratlar pozulyar
	_, err = db.Exec(context.Background(), "DELETE FROM helper_images WHERE image = ANY($1::VARCHAR[])", product.Images)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "data successfully updated",
	})
}

func GetProductByID(c *gin.Context) {
	// initialize database connection
	db, err := config.ConnDB()
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer db.Close()

	// request parametrden category id alynyar
	productID := c.Param("id")

	// database - den request parametr - den gelen id boyunca maglumat cekilyar
	var product models.Product
	if err := db.QueryRow(context.Background(),
		"SELECT id,name_tm,name_ru,name_en,description_tm,description_ru,description_en,category_id FROM products WHERE id = $1", productID).
		Scan(&product.ID, &product.NameTM, &product.NameRU, &product.NameEN,
			&product.DescriptionTM, &product.DescriptionRU, &product.DescriptionEN, &product.CategoryID); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// eger databse sol maglumat yok bolsa error return edilyar
	if product.ID == "" {
		helpers.HandleError(c, 404, "record not found")
		return
	}

	// harydyn suraty db - den alynyar
	rowsImage, err := db.Query(context.Background(), "SELECT image FROM product_images WHERE product_id=$1", productID)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsImage.Close()

	for rowsImage.Next() {
		var image string
		if err := rowsImage.Scan(&image); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		product.Images = append(product.Images, image)
	}

	// harydyn duzumi alynyar
	rowsComposition, err := db.Query(context.Background(),
		"SELECT id,name_tm,name_ru,name_en,percentage FROM product_compositions WHERE product_id=$1", product.ID,
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
			product.Compositions = append(product.Compositions, composition)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"product": product,
	})
}

func GetProducts(c *gin.Context) {
	var requestQuery serializations.CategoryQuery
	var count uint
	var products []models.Product
	deletedAt := `IS NULL`
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

	if requestQuery.IsDeleted {
		deletedAt = `IS NOT NULL`
	}

	if requestQuery.Search != "" {
		incomingsSarch := slug.MakeLang(c.Query("search"), "en")
		search = strings.ReplaceAll(incomingsSarch, "-", " | ")
		searchStr = fmt.Sprintf("%%%s%%", search)
	}

	if requestQuery.Search != "" {
		searchQuery = fmt.Sprintf(` AND (to_tsvector(slug_%s) @@ to_tsquery('%s') OR slug_%s LIKE '%s') `, requestQuery.Lang, search, requestQuery.Lang, searchStr)
	}

	// database - den maglumatlaryn sany alynyar
	queryCount := fmt.Sprintf(`SELECT COUNT(id) FROM products WHERE deleted_at %s %s`, deletedAt, searchQuery)
	if err = db.QueryRow(context.Background(), queryCount).Scan(&count); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database maglumatlar alynyar
	queryProducts := fmt.Sprintf(
		`SELECT id,name_tm,name_ru,name_en,description_tm,description_ru,description_en,category_id FROM products 
		WHERE deleted_at %s %s ORDER BY created_at DESC LIMIT %d OFFSET %d`,
		deletedAt, searchQuery, requestQuery.Limit, offset)
	rowsProduct, err := db.Query(context.Background(), queryProducts)
	if err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}
	defer rowsProduct.Close()

	for rowsProduct.Next() {
		var product models.Product
		if err := rowsProduct.Scan(&product.ID, &product.NameTM, &product.NameRU,
			&product.NameEN, &product.DescriptionTM, &product.DescriptionRU, &product.DescriptionEN, &product.CategoryID); err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}

		//  harydyn suratlary db - den alynyar
		rowsImages, err := db.Query(context.Background(), "SELECT image FROM product_images WHERE product_id=$1", product.ID)
		if err != nil {
			helpers.HandleError(c, 400, err.Error())
			return
		}
		defer rowsImages.Close()

		for rowsImages.Next() {
			var image string
			if err := rowsImages.Scan(&image); err != nil {
				helpers.HandleError(c, 400, err.Error())
				return
			}
			product.Images = append(product.Images, image)
		}

		// harydyn duzumi alynyar
		rowsComposition, err := db.Query(context.Background(),
			"SELECT id,name_tm,name_ru,name_en,percentage FROM product_compositions WHERE product_id=$1", product.ID,
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
				product.Compositions = append(product.Compositions, composition)
			}
		}

		products = append(products, product)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   true,
		"products": products,
		"count":    count,
	})
}

// func DeleteProductByID(c *gin.Context) {
// 	// initialize database connection
// 	db, err := config.ConnDB()
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer db.Close()

// 	// request parametr - den id alynyar
// 	ID := c.Param("id")

// 	// maglumatyn db barlygy barlanyar
// 	if err := helpers.ValidateRecordByID("products", ID, "NULL", db); err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	// eger maglumat bar bolsa onda harydyn suratlary db - den alynyar we pozulyar
// 	rowsImage, err := db.Query(context.Background(), "SELECT image FROM product_images WHERE product_id=$1", ID)
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}
// 	defer rowsImage.Close()

// 	for rowsImage.Next() {
// 		var image string
// 		if err := rowsImage.Scan(&image); err != nil {
// 			helpers.HandleError(c, 400, err.Error())
// 			return
// 		}

// 		// local path - dan surat pozulyar
// 		if err := os.Remove(helpers.ServerPath + image); err != nil {
// 			helpers.HandleError(c, 400, err.Error())
// 			return
// 		}
// 	}

// 	// harydyn suraty pozulandan son haryt we onun bilen baglanysykly maglumatlar pozulyar
// 	_, err = db.Exec(context.Background(), "DELETE FROM products WHERE id = $1", ID)
// 	if err != nil {
// 		helpers.HandleError(c, 400, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"status":  true,
// 		"message": "data successfully deleted",
// 	})
// }
