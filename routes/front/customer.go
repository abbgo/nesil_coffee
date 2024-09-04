package frontApi

import (
	controllers "nesil_coffe/controllers/front"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func CustomerRoutes(back *gin.RouterGroup) {
	api := back.Group("/customers")
	{
		// RegisterCustomer ->  Customer registrasiya etmek ucin ulanylyar.
		api.POST("register", controllers.RegisterCustomer)

		// LoginCustomer Customer login etmek ucin ulanylyar.
		api.POST("login", controllers.LoginCustomer)

		// UpdateCustomer customerin maglumatlaryny uytgetmek ucin ulanylyar.
		api.PUT("update", middlewares.CheckToken(false), controllers.UpdateCustomer)
	}
}
