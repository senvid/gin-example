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
	s.LoadHTMLGlob("templates/*")

	s.GET("/", controllers.Home)
	s.GET("/compose", controllers.ComposeGet)
	s.GET("/login", controllers.LoginGet)
	s.POST("/login", controllers.LoginPost)
	s.GET("/logout", controllers.LogoutGet)
	s.GET("/about", controllers.AboutGet)
	s.GET("/archive", controllers.ArchiveGet)
	// s.POST("/compose", controllers.ComposePost)
	// s.GET("/test/:key", controllers.Test)
	s.Run(":8000")
}
