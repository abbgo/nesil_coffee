package frontApi

import (
	controllers "nesil_coffe/controllers/front"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(back *gin.RouterGroup) {
	api := back.Group("/customers")
	{
		// RegisterCustomer ->  Customer registrasiya etmek ucin ulanylyar.
		api.POST("register", controllers.RegisterCustomer)

		// LoginCustomer Customer login etmek ucin ulanylyar.
		api.POST("login", controllers.LoginCustomer)

		// // UpdateAdmin admin - in maglumatlaryny uytgetmek ucin ulanylyar.
		// api.PUT("update", middlewares.CheckAdmin(), controllers.UpdateAdmin)
	}
}
