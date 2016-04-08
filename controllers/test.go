package controllers
 import (

 	"github.com/gin-gonic/gin"
 	"net/http"
 )

func TestHandler(c *gin.Context)  {
		//c.PostForm("")
	id := c.Param("id")
	if id == "/" {
		id="none"
	}
	c.HTML(http.StatusOK,"index.html",gin.H{
		"id":id,
	})
}


