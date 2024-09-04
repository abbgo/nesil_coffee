package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func SliderBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/sliders").Use(middlewares.CheckToken(true))
	{
		// CreateSlider -> Slider gosmak ulanylar
		api.POST("", controllers.CreateSlider)

		// UpdateSliderByID -> id boyunca Sliderin maglumatlaryny update etmek ucin ulanylyar
		api.PUT("", controllers.UpdateSliderByID)

		// GetSliderByID -> id - si boyunca Sliderin maglumatlaryny almak ucin ulanylyar
		api.GET(":id", controllers.GetSliderByID)

		// GetSliders -> Ahli Sliderlaryn maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetSliders)

		// // DeleteCategoryByID -> id boyunca category pozmak ucin ulanylyar
		// api.DELETE(":id", controllers.DeleteCategoryByID)
	}
}
