package router

import (
	"github.com/gin-gonic/gin"
	"github.com/tesolotech/web-service-gin/controllers"
	"github.com/tesolotech/web-service-gin/middleware"
)

func Routes(route *gin.Engine) {
	album := route.Group("album")
	{
		album.GET("/", middleware.AuthenticateJWT(), controllers.GetAlbums)
		album.GET("/:id", middleware.AuthenticateJWT(), controllers.GetAlbumByID)
		album.POST("/", controllers.PostAlbums)
	}
}
