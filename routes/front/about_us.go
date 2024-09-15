package frontApi

import (
	controllers "nesil_coffe/controllers/back"

	"github.com/gin-gonic/gin"
)

func AboutFrontRoutes(front *gin.RouterGroup) {
	api := front.Group("/about-us")
	{
		// GetOneAboutUs -> 1 sany About Us text - in maglumatlaryny almak ucin ulanylyar
		api.GET("one", controllers.GetOneAboutUs)
	}
}
