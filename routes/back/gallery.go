package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func GalleryBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/galleries").Use(middlewares.CheckToken(true))
	{
		// CreateGallery -> Galareya surat ya-da video gosmak ulanylar
		api.POST("", controllers.CreateGallery)

		// UpdateGalleryByID -> id boyunca galereya maglumatlaryny update etmek ucin ulanylyar
		api.PUT("", controllers.UpdateGalleryByID)

		// GetGalleryByID -> id - si boyunca gallerinin maglumatlaryny almak ucin ulanylyar
		api.GET(":id", controllers.GetGalleryByID)

		// GetGalleries -> Ahli gallery maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetGalleries)

		// DeleteGalleryByID -> id boyunca Product pozmak ucin ulanylyar
		api.DELETE(":id", controllers.DeleteGalleryByID)
	}
}
