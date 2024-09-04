package frontApi

import (
	controllers "nesil_coffe/controllers/front"

	"github.com/gin-gonic/gin"
)

func CategoryFrontRoutes(front *gin.RouterGroup) {
	api := front.Group("/categories")
	{
		// GetCategories -> Ahli Category - leryn maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetCategories)
	}
}
