package controllers

import (
	. "ginsite/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	//"ginsite/config"
)

//var sp int = config.PostsNum
func PageHandler(c *gin.Context) {
	num := c.Param("num")
	onpage, err := strconv.Atoi(num)
	if err != nil {
		c.Redirect(http.StatusBadRequest, "/")
		return
	}
	var posts []Post
	DB.Order("pid desc").Offset(sp).Limit(sp).Find(&posts)
	var countrows int
	DB.Model(&Post{}).Count(&countrows)
	var sumpage int
	if countrows%sp == 0 {
		sumpage = countrows / sp
	} else {
		sumpage = countrows/sp + 1
	}
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"home":    "welcome home",
		"posts":   posts,
		"sumpage": sumpage,
		"onpage":  onpage,
	})

}
