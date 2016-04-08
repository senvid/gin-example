package controllers

 import (
 	"github.com/gin-gonic/gin"
 	"net/http"
 )




func AboutHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", gin.H{})
}


