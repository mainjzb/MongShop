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
func Routers() *gin.Engine {
	r := gin.Default()
	r.StaticFS("/static", http.Dir("/static"))
	//Routers.Use(middleware.LoadTls())
	log.Println("use middleware cors")
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	log.Println("register swagger handler")
	v1 := r.Group("/api/v1")
	{
		v1.GET("/message", service.GetMessage)
		v1.GET("/hello", service.Hello)
	}
	r.Run(":8080")
	return r
}
