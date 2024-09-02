package routes

import (
	backApi "nesil_coffe/routes/back"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	routes := gin.Default()

	routes.Use(gzip.Gzip(gzip.DefaultCompression))

	routes.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "RefreshToken", "Authorization"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	back := routes.Group("/api/back")
	{
		// bu route - ler back kategoriyalar ucin doredilen rout - laryn toplumy
		backApi.CategoryBackRoutes(back)

		// bu route - ler back suratlar ucin doredilen rout - laryn toplumy
		backApi.BackImageRoutes(back)

		// bu route - ler back harytlar ucin doredilen rout - laryn toplumy
		backApi.ProductBackRoutes(back)

		// bu route - ler back gallery ucin doredilen rout - laryn toplumy
		backApi.GalleryBackRoutes(back)
	}

	return routes
}
