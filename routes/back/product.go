package back

import (
	controllers "nesil_coffe/controllers/back"

	"github.com/gin-gonic/gin"
)

func ProductBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/products")
	{
		// CreateProduct -> Product gosmak ulanylar
		api.POST("", controllers.CreateProduct)

		// UpdateProductByID -> id boyunca Product - in maglumatlaryny update etmek ucin ulanylyar
		api.PUT("", controllers.UpdateProductByID)

		// GetBrendBGetProductByIDyID -> id - si boyunca Product - in maglumatlaryny almak ucin ulanylyar
		api.GET(":id", controllers.GetProductByID)

		// GetProducts -> Ahli Product - laryn maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetProducts)

		// DeleteCategoryByID -> id boyunca category pozmak ucin ulanylyar
		api.DELETE(":id", controllers.DeleteCategoryByID)
	}
}
