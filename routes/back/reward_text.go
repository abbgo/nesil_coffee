package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func RewardTextBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/reward-text").Use(middlewares.CheckToken(true))
	{
		// CreateRewrdText -> Reward Text gosmak ulanylar
		api.POST("", controllers.CreateRewardText)

		// UpdateRewrdTextByID -> id boyunca Reward Text - in maglumatlaryny update etmek ucin ulanylyar
		api.PUT("", controllers.UpdateRewardTextByID)

		// GetOneRewrdText -> 1 sany Reward Text - in maglumatlaryny almak ucin ulanylyar
		api.GET("one", controllers.GetOneRewardText)

		// DeleteRewrdText -> Reward Text pozmak ucin ulanylyar
		api.DELETE(":id", controllers.DeleteRewardText)
	}
}
