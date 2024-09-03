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
		admin.PUT("update", middlewares.CheckAdmin(), controllers.UpdateAdmin)

		// // UpdateAdminPassword admin - in parolyny uytgetmek ucin ulanylyar.
		// admin.PUT("update-password", middlewares.IsSuperAdmin(), controllers.UpdateAdminPassword)

		// // GetAdmin funksiya haeder - den gelen id boyunca bir sany admin - i almak ucin ulanylyar.
		// admin.GET("one", middlewares.CheckToken("admin"), controllers.GetAdmin)

		// // GetAdmins funksiya hemme admin - leri almak ucin ulanylyar.
		// admin.GET("", middlewares.IsSuperAdmin(), controllers.GetAdmins)

		// // DeletePermanentlyAdminByID funksiya id boyunca admin - i doly ( korzinadan ) pozmak ucin ulanylyar
		// admin.DELETE(":id/delete", middlewares.IsSuperAdmin(), controllers.DeletePermanentlyAdminByID)
	}
}
