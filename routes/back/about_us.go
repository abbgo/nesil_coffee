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

		// UpdateTextSliderByID -> id boyunca Text Slider - in maglumatlaryny update etmek ucin ulanylyar
		api.PUT("", controllers.UpdateTextSliderByID)

		// GetOneTextSlider -> 1 sany Text Slider - in maglumatlaryny almak ucin ulanylyar
		api.GET("one", controllers.GetOneTextSlider)

		// DeleteTextSlider -> Text SLider pozmak ucin ulanylyar
		api.DELETE(":id", controllers.DeleteTextSlider)
	}
}
