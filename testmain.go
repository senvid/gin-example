package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/gin-gonic/gin/binding"
	// "ginsite/middleware"
	"ginsite/models"
	// "log"
	"net/http"
	// "time"
)

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	models.InitDB()
	r := gin.Default()
	r.Static("/static", "./static")
	r.StaticFile("/favicon.ico", "./static/favicon.ico")
	r.LoadHTMLGlob("templates/*")

	// r.Use(middleware.Mytime())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "welcome home",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(200, "hello %s", name)

	})

	// curl  --form user=u1 --form password=p1 localhost:8000/login
	var form LoginForm
	r.POST("/login", func(c *gin.Context) {
		if c.Bind(&form) == nil {
			if form.User == "u1" && form.Password == "p1" {
				c.JSON(200, gin.H{"status": "success to login"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		}
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// You also can use a struct
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// Note that msg.Name becomes "user" in the JSON
		// Will output  :   {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "main website",
		})
	})

	r.GET("/redic", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	// authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
	// 	"foo":    "bar",
	// 	"austin": "1234",
	// 	"lena":   "hello2",
	// 	"manu":   "4321",
	// }))

	// // /admin/secrets endpoint
	// // hit "localhost:8080/admin/secrets
	// authorized.GET("/secrets", func(c *gin.Context) {
	// 	// get user, it was set by the BasicAuth middleware
	// 	user := c.MustGet(gin.AuthUserKey).(string)
	// 	if secret, ok := secrets[user]; ok {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	// 	}
	// })

	gin.SetMode(gin.DebugMode)
	r.Run(":8000")
}
