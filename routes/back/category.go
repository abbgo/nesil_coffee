package back

import (
	controllers "nesil_coffe/controllers/back"

	"github.com/gin-gonic/gin"
)

func CategoryBackRoutes(back *gin.RouterGroup) {
	backCategoryApi := back.Group("/categories")
	{
		// CreateCategory -> Category gosmak ulanylar
		backCategoryApi.POST("", controllers.CreateCategory)

		// UpdateCategoryByID -> id boyunca Category - in maglumatlaryny update etmek ucin ulanylyar
		backCategoryApi.PUT("", controllers.UpdateCategoryByID)

		// GetBrendByID -> id - si boyunca Category - in maglumatlaryny almak ucin ulanylyar
		backCategoryApi.GET(":id", controllers.GetCategoryByID)

		// GetCategories -> Ahli Category - leryn maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		backCategoryApi.GET("", controllers.GetCategories)

		// DeleteCategoryByID -> id boyunca category pozmak ucin ulanylyar
		backCategoryApi.DELETE(":id", controllers.DeleteCategoryByID)
	}
}
