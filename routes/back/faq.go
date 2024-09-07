package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func FAQBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/faqs").Use(middlewares.CheckToken(true))
	{
		// CreateFAQ -> FAQ gosmak ulanylar
		api.POST("", controllers.CreateFAQ)

		// UpdateFAQByID -> id boyunca FAQ - nin maglumatlaryny update etmek ucin ulanylyar
		api.PUT("", controllers.UpdateFAQByID)

		// GetFAQByID -> id - si boyunca FAQ - nin maglumatlaryny almak ucin ulanylyar
		api.GET(":id", controllers.GetFAQByID)

		// GetFAQs -> Ahli FAQ - laryn maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetFAQs)

		// DeleteFAQByID -> id boyunca FAQ -ny pozmak ucin ulanylyar
		api.DELETE(":id", controllers.DeleteFAQByID)
	}
}
