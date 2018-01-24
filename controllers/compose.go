package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ComposeHandler(c *gin.Context) {
	islogin, _ = c.Get("islogin")
	c.HTML(http.StatusOK, "compose.tmpl", gin.H{
		"islogin": islogin,
	})
}
