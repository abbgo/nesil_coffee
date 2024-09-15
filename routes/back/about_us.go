package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func AboutBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/about-us").Use(middlewares.CheckToken(true))
	{
		// CreateAboutUs ->About Us text gosmak ulanylar
		api.POST("", controllers.CreateAboutUs)

		// UpdateAboutUsByID -> id boyunca About us text - in maglumatlaryny update etmek ucin ulanylyar
		api.PUT("", controllers.UpdateAboutUsByID)

		// GetOneAboutUs -> 1 sany About Us text - in maglumatlaryny almak ucin ulanylyar
		api.GET("one", controllers.GetOneAboutUs)

		// DeleteAboutUs -> About Us text pozmak ucin ulanylyar
		api.DELETE(":id", controllers.DeleteAboutUs)
	}
}
