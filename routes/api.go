package routes

import (
	adminApi "nesil_coffe/routes/admin"
	backApi "nesil_coffe/routes/back"
	frontApi "nesil_coffe/routes/front"

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

		// bu route - ler back video ucin doredilen rout - laryn toplumy
		backApi.BackVideoRoutes(back)

		// bu route - ler back resepteler ucin doredilen rout - laryn toplumy
		backApi.RecipeBackRoutes(back)

		// bu route - ler back slider ucin doredilen rout - laryn toplumy
		backApi.SliderBackRoutes(back)

		// bu route - ler back mailler ucin doredilen rout - laryn toplumy
		backApi.MailBackRoutes(back)

		// bu route - ler back faq ucin doredilen rout - laryn toplumy
		backApi.FAQBackRoutes(back)
	}

	admin := routes.Group("/api")
	{
		// bu route - ler admin - ler ucin doredilen route - laryn toplumy
		adminApi.AdminRoutes(admin)
	}

	front := routes.Group("/api")
	{
		// bu route - ler klient - ler ucin doredilen route - laryn toplumy
		frontApi.CustomerRoutes(front)

		// bu route - ler harytlar ucin doredilen route - laryn toplumy
		frontApi.ProductFrontRoutes(front)

		// bu route - ler kategoriyalar ucin doredilen route - laryn toplumy
		frontApi.CategoryFrontRoutes(front)

		// bu route - ler gallery ucin doredilen route - laryn toplumy
		frontApi.GalleryFrontRoutes(front)

		// bu route - ler gallery ucin doredilen route - laryn toplumy
		frontApi.RecipeFrontRoutes(front)

		// bu route - ler slider ucin doredilen route - laryn toplumy
		frontApi.SliderFrontRoutes(front)

		// bu route - ler mail ucin doredilen route - laryn toplumy
		frontApi.MailFrontRoutes(front)
	}

	return routes
}
