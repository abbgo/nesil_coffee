package frontApi

import (
	controllers "nesil_coffe/controllers/back"

	"github.com/gin-gonic/gin"
)

func RewardTextFrontRoutes(back *gin.RouterGroup) {
	api := back.Group("/reward-text")
	{
		// GetOneRewrdText -> 1 sany Reward Text - in maglumatlaryny almak ucin ulanylyar
		api.GET("one", controllers.GetOneRewardText)
	}
}
