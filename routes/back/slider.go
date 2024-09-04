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

		// // GetCategoryByID -> id - si boyunca Category - in maglumatlaryny almak ucin ulanylyar
		// api.GET(":id", controllers.GetCategoryByID)

		// // GetCategories -> Ahli Category - leryn maglumatlaryny request query - den gelen
		// // limit we page boyunca pagination ulanyp almak ucin ulanylyar
		// api.GET("", controllers.GetCategories)

		// // DeleteCategoryByID -> id boyunca category pozmak ucin ulanylyar
		// api.DELETE(":id", controllers.DeleteCategoryByID)
	}
}
