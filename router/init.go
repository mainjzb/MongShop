package router

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	_ "mongShop/docs"
	"mongShop/service"
	"net/http"
)

// @title goweb project
// @version 1.0
// @description this is goweb server.
// @host localhost:8080
// @BasePath /api/v1
func Router() *gin.Engine {
	Router := gin.Default()
	Router.StaticFS("/static", http.Dir("/static"))
	//Router.Use(middleware.LoadTls())
	log.Println("use middleware cors")
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Println("register swagger handler")
	v1 := Router.Group("/api/v1")
	{
		v1.GET("/message", service.GetMessage)
		v1.GET("/hello", service.Hello)
	}
	Router.Run(":8080")
	return Router
}
