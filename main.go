package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tesolotech/web-service-gin/router"
)

func main() {
	fmt.Println("Welcome to web service with gin")
	routers := gin.Default()
	router.Routes(routers)
	routers.Group("album")
	routers.Run("localhost:8080")
}

/* Routes : GET - http://localhost:8080/album {Header Token key : 'Authorization' : "Get From post api response"}
			POST : Create album payload :
{
    "title" : "Ye Kaali Kaali Aankhen",
 	"artist" : "Kumar Sonu",
 	"price" : 500
}
*/
