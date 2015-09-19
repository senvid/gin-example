package main

import (
	"gin-example/controllers"
	"gin-example/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.InitDB()
	s := gin.Default()
	s.Static("/static", "./static")
	s.LoadHTMLGlob("template/*")
	s.GET("/", controllers.Mydb)
	s.GET("/s", controllers.Hello)
	s.GET("/test/:key", controllers.Hello)
	s.Run(":8000")
}
