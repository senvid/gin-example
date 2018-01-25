package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AboutHandler(c *gin.Context) {
	islogin, _ := c.Get("islogin")
	c.HTML(http.StatusOK, "about.html", gin.H{
		"islogin": islogin,
		"ok":      "2",
	})
}
