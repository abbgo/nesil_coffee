package frontApi

import (
	controllers "nesil_coffe/controllers/front"

	"github.com/gin-gonic/gin"
)

func MailFrontRoutes(front *gin.RouterGroup) {
	api := front.Group("/mails")
	{
		// SendMail -> firma mail ugratmak ucin ulanylyar
		api.POST("send-mail", controllers.SendMail)
	}
}
