package frontApi

import (
	controllers "nesil_coffe/controllers/front"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(back *gin.RouterGroup) {
	api := back.Group("/customers")
	{
		// RegisterAdmin admin - i registrasiya etmek ucin ulanylyar.
		api.POST("register", controllers.RegisterCustomer)

		// // LoginAdmin admin - i login etmek ucin ulanylyar.
		// api.POST("login", controllers.LoginAdmin)

		// // UpdateAdmin admin - in maglumatlaryny uytgetmek ucin ulanylyar.
		// api.PUT("update", middlewares.CheckAdmin(), controllers.UpdateAdmin)
	}
}
