package controllers

 import (
 	"github.com/gin-gonic/gin"
 	"net/http"
 )

func ComposeHandler(c *gin.Context) {

	c.HTML(http.StatusOK, "compose.tmpl", gin.H{})
}
