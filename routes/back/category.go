package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoryBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/categories").Use(middlewares.CheckToken(true))
	{
		// CreateCategory -> Category gosmak ulanylar
		api.POST("", controllers.CreateCategory)

		// UpdateCategoryByID -> id boyunca Category - in maglumatlaryny update etmek ucin ulanylyar
		api.PUT("", controllers.UpdateCategoryByID)

		// GetCategoryByID -> id - si boyunca Category - in maglumatlaryny almak ucin ulanylyar
		api.GET(":id", controllers.GetCategoryByID)

		// GetCategories -> Ahli Category - leryn maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetCategories)

		// DeleteCategoryByID -> id boyunca category pozmak ucin ulanylyar
		api.DELETE(":id", controllers.DeleteCategoryByID)
	}
}
