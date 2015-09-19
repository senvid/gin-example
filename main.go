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
	// s.LoadHTMLGlob("templates/*")
	s.LoadHTMLFiles("templates/home.html", "templates/compose.html", "templates/404.html")

	s.GET("/", controllers.Home)
	s.GET("/compose", controllers.ComposeGet)
	s.POST("/compose", controllers.ComposePost)
	// s.GET("/test/:key", controllers.Test)
	s.Run(":8000")
}