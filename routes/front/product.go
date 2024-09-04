package frontApi

import (
	controllers "nesil_coffe/controllers/front"

	"github.com/gin-gonic/gin"
)

func ProductFrontRoutes(front *gin.RouterGroup) {
	api := front.Group("/products")
	{
		// GetProductByID -> id - si boyunca Product - in maglumatlaryny almak ucin ulanylyar
		api.GET(":id", controllers.GetProductByID)

		// GetProducts -> Ahli Product - laryn maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetProducts)
	}
}
