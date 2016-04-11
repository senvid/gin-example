package controllers

import (
	"ginsite/component"
	. "ginsite/models"
	"ginsite/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func GetAllHandler(c *gin.Context) {

	var posts []Post
	DB.Select("title,slug,published").Order("pid desc").Find(&posts)
	titles := component.GetTitle()
	tags := component.GetTag()

	c.HTML(http.StatusOK, "posts.html", gin.H{
		"posts":  posts,
		"titles": titles,
		"tags":   tags,
	})
}

func NewPostAndEditHandler(c *gin.Context) {
	//  post: --> /post/   or  -->/post/*slug

	Title := c.PostForm("title")
	Content := c.PostForm("content")
	Published := time.Now()
	now := time.Now()

	slug := c.PostForm("slug")
	//  /post/*slug  ---> slug start with "/"
	var Slug string
	if slug == "/" {
		Slug = now.Format("2006/1/02/") + strconv.Itoa(now.Nanosecond())
	} else {
		Slug = string([]byte(slug)[1:])
	}

	uid, err := util.GetSecureCookie(c.Request, "user")
	if err != nil {
		c.Redirect(http.StatusNonAuthoritativeInfo, "/login")
		return
	}

	Useridstr, _ := strconv.Atoi(uid)
	Userid := uint(Useridstr)

	tag := c.PostForm("tag")
	var (
		tags  Tag
		Tagid uint = 0
	)
	if tag != "" {
		DB.Where("type=?", tag).Find(&tags)
		Tagid = tags.Tid
	}

	post := Post{
		Title:     Title,
		Content:   Content,
		Published: Published,
		Slug:      Slug,
		Userid:    Userid,
		Tagid:     Tagid,
	}
	if slug == "/" {
		DB.Create(&post)
	} else {
		DB.Model(&post).Update("title", "content", "slug", "tagid")
	}

	c.Redirect(http.StatusMovedPermanently, "/posts/"+Slug)
}

func DeleteHandler(c *gin.Context) {
	slug := c.Query("slug")
	if slug != "" {
		DB.Delete(&Post{}).Where("slug = ?", slug)
		c.JSON(http.StatusOK, gin.H{
			"slug":   slug,
			"status": "ok",
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"slug":   slug,
		"status": "fail",
	})
}

func GetOneHandler(c *gin.Context) {
	slug := c.Param("slug")
	if slug == "" {
		c.Redirect(http.StatusMovedPermanently, "/")
	}
	var posts []Post
	DB.Where("slug = ?", slug).Find(&posts)
	titles := component.GetTitle()
	tags := component.GetTag()
	c.HTML(http.StatusOK, "posts.html", gin.H{
		"posts":  posts,
		"titles": titles,
		"tags":   tags,
	})

}
