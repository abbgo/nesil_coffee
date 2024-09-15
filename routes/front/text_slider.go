package frontApi

import (
	controllers "nesil_coffe/controllers/back"

	"github.com/gin-gonic/gin"
)

func TextSliderFrontRoutes(front *gin.RouterGroup) {
	api := front.Group("/text-slider")
	{
		// GetOneTextSlider -> 1 sany Text Slider - in maglumatlaryny almak ucin ulanylyar
		api.GET("one", controllers.GetOneTextSlider)
	}
}
