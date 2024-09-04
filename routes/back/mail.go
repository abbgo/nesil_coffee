package back

import (
	controllers "nesil_coffe/controllers/back"
	"nesil_coffe/middlewares"

	"github.com/gin-gonic/gin"
)

func MailBackRoutes(back *gin.RouterGroup) {
	api := back.Group("/mails").Use(middlewares.CheckToken(true))
	{
		api.GET("", controllers.GetMails)

		// // DeleteGalleryByID -> id boyunca Product pozmak ucin ulanylyar
		// api.DELETE(":id", controllers.DeleteGalleryByID)
	}
}
