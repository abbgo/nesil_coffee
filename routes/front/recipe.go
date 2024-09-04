package frontApi

import (
	controllers "nesil_coffe/controllers/front"

	"github.com/gin-gonic/gin"
)

func RecipeFrontRoutes(front *gin.RouterGroup) {
	api := front.Group("/recipes")
	{
		// GetRecipeByID -> id - si boyunca Reseptin maglumatlaryny almak ucin ulanylyar
		api.GET(":id", controllers.GetRecipeByID)

		// GetRecipes -> Ahli Reseptelerin maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetRecipes)
	}
}
