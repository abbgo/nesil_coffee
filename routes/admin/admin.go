package adminApi

import (
	controllers "nesil_coffe/controllers/admin"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(back *gin.RouterGroup) {
	admin := back.Group("/admins")
	{
		// RegisterAdmin admin - i registrasiya etmek ucin ulanylyar.
		admin.POST("register", controllers.RegisterAdmin)

		// LoginAdmin admin - i login etmek ucin ulanylyar.
		admin.POST("login", controllers.LoginAdmin)

		// UpdateAdmin admin - in maglumatlaryny uytgetmek ucin ulanylyar.
		admin.PUT("update", middlewares.CheckToken(true), controllers.UpdateAdmin)
	}
}
