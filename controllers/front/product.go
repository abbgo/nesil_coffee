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

func GetProducts(c *gin.Context) {
	var requestQuery serializations.CategoryQuery
	var count uint
	var products []models.Product
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
	queryCount := fmt.Sprintf(`SELECT COUNT(id) FROM products %s `, searchQuery)
	if err = db.QueryRow(context.Background(), queryCount).Scan(&count); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database maglumatlar alynyar
	queryProducts := fmt.Sprintf(
		`SELECT id,name_tm,name_ru,name_en,description_tm,description_ru,description_en,category_id FROM products 
		%s ORDER BY created_at DESC LIMIT %d OFFSET %d`,
		searchQuery, requestQuery.Limit, offset)
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
