package frontApi

import (
	controllers "nesil_coffe/controllers/front"

	"github.com/gin-gonic/gin"
)

func GalleryFrontRoutes(front *gin.RouterGroup) {
	api := front.Group("/galleries")
	{
		// GetGalleries -> Ahli gallery maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetGalleries)
	}
}
