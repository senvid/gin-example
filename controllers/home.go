package controllers

import (
	. "ginsite/models"
	"github.com/gin-gonic/gin"
	"net/http"

	"ginsite/config"

	"ginsite/component"
)

var sp int = config.PostsNum

func HomeHandler(c *gin.Context) {
	// "SELECT * FROM posts ORDER BY id DESC LIMIT %s", sp
	// "SELECT count(*) FROM posts"
	var posts []Post
	DB.Order("pid desc").Limit(sp).Find(&posts)

	if len(posts) == 0 {
		c.Redirect(http.StatusFound, "/user/compose")
		return
	}
	var countrows int
	DB.Model(&Post{}).Count(&countrows)

	var sumpage int
	if countrows%sp == 0 {
		sumpage = countrows / sp
	} else {
		sumpage = countrows/sp + 1
	}
	//title and tag
	titles := component.GetTitle()
	tags := component.GetTag()
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"home":    "welcome home",
		"posts":   posts,
		"sumpage": sumpage,
		"onpage":  1,
		"current": c.MustGet("current").(bool),
		"titles":  titles,
		"tags":    tags,
	})
}
