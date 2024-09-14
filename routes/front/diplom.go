package frontApi

import (
	controllers "nesil_coffe/controllers/front"

	"github.com/gin-gonic/gin"
)

func DiplomFrontRoutes(front *gin.RouterGroup) {
	api := front.Group("/diploms")
	{
		// GetDiploms -> Ahli diplomlaryn maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetDiploms)
	}
}
