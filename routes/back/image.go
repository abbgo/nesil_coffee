package back

import (
	controllers "nesil_coffe/controllers/back"

	"github.com/gin-gonic/gin"
)

func BackImageRoutes(back *gin.RouterGroup) {
	backImageApi := back.Group("/images")
	{
		// surat gosmak ya-da uytgetmek ucin ulanylyar
		backImageApi.POST("", controllers.AddOrUpdateImage)

		// // surat pozmak ucin ulanylyar
		backImageApi.DELETE("", controllers.DeleteImage)
	}
}
