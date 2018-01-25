package main

import (
	"ginsite/controllers"
	"ginsite/middleware"
	"ginsite/models"
	"github.com/gin-gonic/gin"
	"runtime"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	models.InitDB()

	r := gin.Default()
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLGlob("templates/**/*")

	r.Use(middleware.Current())

	authorized := r.Group("/")
	authorized.Use(middleware.AuthRequired())
	{
		authorized.POST("/posts/*slug", controllers.NewPostAndEditHandler)
		authorized.DELETE("/posts/:slug", controllers.DeleteHandler)
	}
	r.GET("/posts", controllers.GetAllHandler)
	r.GET("/posts/:slug", controllers.GetOneHandler)

	r.GET("/", controllers.HomeHandler)
	r.GET("/compose", controllers.ComposeHandler)
	r.GET("/aside", controllers.AsideHandler)

	r.GET("/page/:num", controllers.PageHandler)

	r.GET("/login", controllers.LoginHandler)
	r.POST("/login", controllers.LoginPostHandler)
	r.GET("/logout", controllers.LogoutHandler)

	r.GET("/about", controllers.AboutHandler)

	r.GET("/test/*id", controllers.TestHandler)

	// r.PUT("/posts/:id", ...)
	// r.DELETE("/posts/:id", ...)
	//r.GET("*",models.PageNotFoundHandler)

	// GET /zoos：列出所有动物园
	// POST /zoos：新建一个动物园
	// GET /zoos/ID：获取某个指定动物园的信息
	// PUT /zoos/ID：更新某个指定动物园的信息（提供该动物园的全部信息）
	// PATCH /zoos/ID：更新某个指定动物园的信息（提供该动物园的部分信息）
	// DELETE /zoos/ID：删除某个动物园
	// GET /zoos/ID/animals：列出某个指定动物园的所有动物
	// DELETE /zoos/ID/animals/ID：删除某个指定动物园的指定动物

	// ?limit=10：指定返回记录的数量
	// ?offset=10：指定返回记录的开始位置。
	// ?page=2&per_page=100：指定第几页，以及每页的记录数。
	// ?sortby=name&order=asc：指定返回结果按照哪个属性排序，以及排序顺序。
	// ?animal_type_id=1：指定筛选条件

	gin.SetMode(gin.DebugMode)
	r.Run(":8000")
}
