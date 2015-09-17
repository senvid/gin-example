package controllers

import (
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	r := c.Param("key")
	// c.String(http.StatusOK, r)
	c.HTML(200, "home.html", gin.H{
		"args": r,
	})

}
