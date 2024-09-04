package frontApi

import (
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func CategoryFrontRoutes(front *gin.RouterGroup) {
	api := front.Group("/categories").Use(middlewares.CheckToken(true))
	{
		// GetCategories -> Ahli Category - leryn maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		// api.GET("", controllers.GetCategories)
	}
}
