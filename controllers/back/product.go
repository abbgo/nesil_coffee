package controllers

import (
	"context"
	"nesil_coffe/config"
	"nesil_coffe/helpers"
	"nesil_coffe/models"
	"net/http"

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

	if len(product.Images) == 0 {
		helpers.HandleError(c, 400, "images of product is requred")
		return
	}

	// eger maglumatlar dogry bolsa db maglumatlar gosulyar we gosulandan son gosulan maglumatyn id - si return edilyar
	var productID string
	if err := db.QueryRow(context.Background(),
		"INSERT INTO products (name,description,slug) VALUES ($1,$2,$3) RETURNING id",
		product.Name, product.Description, slug.MakeLang(product.Name, "en"),
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

	// bind edilen maglumatlar barlanyar
	if product.ID == "" {
		helpers.HandleError(c, 400, "product id is required")
		return
	}
	if err := helpers.ValidateRecordByID("products", product.ID, "NULL", db); err != nil {
		helpers.HandleError(c, 400, err.Error())
		return
	}

	// database - daki maglumatlary request body - dan gelen maglumatlar bilen calysyas
	_, err = db.Exec(context.Background(),
		"UPDATE products SET name=$1 , description=$2 , slug=$3 WHERE id=$4",
		product.Name, product.Description, slug.MakeLang(product.Name, "en"), product.ID,
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
