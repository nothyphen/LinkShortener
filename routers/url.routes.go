package routers

import (
	v1 "linkshortner/api/v1"
	"linkshortner/database"
	"linkshortner/middleware"
	"linkshortner/repository"
	"linkshortner/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	postdb			    *gorm.DB						= database.ConnectPostgres()
	linkrepository		repository.LinkRepository 		= repository.NewLinkRepository(postdb)
	linkservices		services.LinkService			= services.NewLinkService(linkrepository)
	shortService		services.ShortService			= services.NewShortRepository(linkrepository)
	linkAPI				v1.LinkApi						= v1.NewLinkApi(linkservices)
	shortAPI			v1.ShortAPI						= v1.NewShortAPI(shortService)
)

func Urls() *gin.Engine {

	route := gin.Default()
	route.Use(middleware.CORSMiddleware())
	route.NoRoute(middleware.NoRouteHandler())
	route.NoMethod(middleware.NoMethodHandler())

	apiV1 := route.Group("api/v1")
	{
		link := apiV1.Group("link")
		{
			link.GET("/")
			link.POST("/", linkAPI.Link)
		}
		short := apiV1.Group("short")
		{
			short.POST("/", shortAPI.Short)
		}
	}

	return route
}