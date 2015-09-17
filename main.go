package main

import (
	. "gin-example/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	s := gin.Default()
	s.Static("/static", "./static")
	s.LoadHTMLGlob("template/*")
	s.GET("/", Hello)
	s.GET("/test/:key", Hello)
	s.Run(":8000")

}
