package frontApi

import (
	controllers "nesil_coffe/controllers/front"

	"github.com/gin-gonic/gin"
)

func FAQFrontRoutes(front *gin.RouterGroup) {
	api := front.Group("/faqs")
	{
		// GetFAQs -> Ahli FAQ - laryn maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetFAQs)
	}
}
