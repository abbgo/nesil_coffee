package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func TextSliderBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/text-slider").Use(middlewares.CheckToken(true))
	{
		// CreateTextSlider -> Text Slider gosmak ulanylar
		api.POST("", controllers.CreateTextSlider)

		// UpdateTextSliderByID -> id boyunca Text Slider - in maglumatlaryny update etmek ucin ulanylyar
		api.PUT("", controllers.UpdateTextSliderByID)

		// GetOneTextSlider -> 1 sany Text Slider - in maglumatlaryny almak ucin ulanylyar
		api.GET("one", controllers.GetOneTextSlider)

		// // GetProducts -> Ahli Product - laryn maglumatlaryny request query - den gelen
		// // limit we page boyunca pagination ulanyp almak ucin ulanylyar
		// api.GET("", controllers.GetProducts)

		// // DeleteProductByID -> id boyunca Product pozmak ucin ulanylyar
		// api.DELETE(":id", controllers.DeleteProductByID)
	}
}
