package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func BackImageRoutes(back *gin.RouterGroup) {
	backImageApi := back.Group("/images")
	{
		// surat gosmak ya-da uytgetmek ucin ulanylyar
		backImageApi.POST("", controllers.AddOrUpdateImage).Use(middlewares.CheckToken(true))

		// // surat pozmak ucin ulanylyar
		backImageApi.DELETE("", controllers.DeleteImage)
	}
}
