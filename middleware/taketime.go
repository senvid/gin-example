package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Mytime() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		// c.Set("example", "12345")

		// before request
		if c.HandlerName() == "ginsite/models.TestHandler"{
			c.Abort()
			c.Redirect(301,"/about")
			log.Println("i want abort.1............")
			return
		}

		c.Next()
		log.Println("i want abort.2............")
		// after request
		latency := time.Since(t)
		log.Print(latency)

		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)

	}
}
