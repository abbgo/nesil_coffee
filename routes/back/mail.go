package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func MailBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/mails").Use(middlewares.CheckToken(true))
	{
		// GetMails -> back mailleri almak ucin ulanylyar
		// pagination ucin request query - de limit we page gelyar
		api.GET("", controllers.GetMails)

		// DeleteMailByID -> id boyunca mail pozmak ucin ulanylyar
		api.DELETE(":id", controllers.DeleteMailByID)
	}
}
