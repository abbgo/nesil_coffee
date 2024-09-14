package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func DiplomBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/diploms").Use(middlewares.CheckToken(true))
	{
		// CreateDiplom -> Diplom gosmak ulanylar
		api.POST("", controllers.CreateDiplom)

		// UpdateDiplomByID -> id boyunca diplom maglumatlaryny update etmek ucin ulanylyar
		api.PUT("", controllers.UpdateDiplomByID)

		// GetDiplomByID -> id - si boyunca diplomyn maglumatlaryny almak ucin ulanylyar
		api.GET(":id", controllers.GetDiplomByID)

		// // GetGalleries -> Ahli gallery maglumatlaryny request query - den gelen
		// // limit we page boyunca pagination ulanyp almak ucin ulanylyar
		// api.GET("", controllers.GetGalleries)

		// // DeleteGalleryByID -> id boyunca Product pozmak ucin ulanylyar
		// api.DELETE(":id", controllers.DeleteGalleryByID)
	}
}
