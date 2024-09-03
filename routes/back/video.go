package back

import (
	controllers "nesil_coffe/controllers/back"

	"github.com/gin-gonic/gin"
)

func BackVideoRoutes(back *gin.RouterGroup) {
	api := back.Group("/videos")
	{
		// video gosmak ya-da uytgetmek ucin ulanylyar
		api.POST("", controllers.AddOrUpdateVideo)

	}
}
