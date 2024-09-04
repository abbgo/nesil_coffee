package frontApi

import (
	controllers "nesil_coffe/controllers/front"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(front *gin.RouterGroup) {
	api := front.Group("/customers")
	{
		// RegisterCustomer ->  Customer registrasiya etmek ucin ulanylyar.
		api.POST("register", controllers.RegisterCustomer)

		// LoginCustomer Customer login etmek ucin ulanylyar.
		api.POST("login", controllers.LoginCustomer)

		// GetCustomer customerin maglumatlaryny almak ucin ulanylyar.
		api.GET("one", middlewares.CheckToken(false), controllers.GetCustomer)

		// UpdateCustomer customerin maglumatlaryny uytgetmek ucin ulanylyar.
		api.PUT("update", middlewares.CheckToken(false), controllers.UpdateCustomer)
	}
}
