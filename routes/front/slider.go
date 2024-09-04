package frontApi

import (
	controllers "nesil_coffe/controllers/front"

	"github.com/gin-gonic/gin"
)

func SliderFrontRoutes(front *gin.RouterGroup) {
	api := front.Group("/sliders")
	{
		// GetSliders -> Ahli Sliderlaryn maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetSliders)
	}
}
