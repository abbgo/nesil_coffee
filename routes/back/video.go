package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func BackVideoRoutes(back *gin.RouterGroup) {
	api := back.Group("/videos").Use(middlewares.CheckToken(true))
	{
		// video gosmak ya-da uytgetmek ucin ulanylyar
		api.POST("", controllers.AddOrUpdateVideo)
	}
}
