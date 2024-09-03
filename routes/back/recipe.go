package back

import (
	controllers "nesil_coffe/controllers/back"

	"github.com/gin-gonic/gin"
)

func RecipeBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/recipes")
	{
		// CreateRecipe -> Resepte gosmak ulanylar
		api.POST("", controllers.CreateRecipe)

		// UpdateRecipeByID -> id boyunca Resepte - in maglumatlaryny update etmek ucin ulanylyar
		api.PUT("", controllers.UpdateRecipeByID)

		// GetRecipeByID -> id - si boyunca Reseptin maglumatlaryny almak ucin ulanylyar
		api.GET(":id", controllers.GetRecipeByID)

		// GetRecipes -> Ahli Reseptelerin maglumatlaryny request query - den gelen
		// limit we page boyunca pagination ulanyp almak ucin ulanylyar
		api.GET("", controllers.GetRecipes)

		// DeleteProductByID -> id boyunca Product pozmak ucin ulanylyar
		api.DELETE(":id", controllers.DeleteProductByID)
	}
}
