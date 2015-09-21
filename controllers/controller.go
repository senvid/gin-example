package controllers

import (
	"crypto/rand"
	"encoding/base64"
	"gin-example/models"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
	// "time"
)

/*
func Test(c *gin.Context) {
	// /test/:key
	r := c.Param("key")
	// c.String(http.StatusOK, r)
	c.HTML(200, "home.html", gin.H{
		"args": r,
	})
}

func Args(c *gin.Context) {
	// /test?key=123
	// r := c.Request.URL.Query().Get("key")
	r := c.Query("key")
	c.HTML(http.StatusOK, "home.html", gin.H{
		"args": r,
	})
}

func Mydb(c *gin.Context) {
	type Result struct {
		Uid   uint
		Email string
	}
	var Res Result
	// u := &models.Users{Email: "test@test.com", Password: "test", Nickname: "test"}
	// models.Mdb.NewRecord(u)
	// models.Mdb.Create(u)
	models.DB.Raw("SELECT uid,email FROM users where uid=1").Scan(&Res)
	c.HTML(200, "home.html", gin.H{
		"args": Res,
	})
}
*/

const (
	sp int = 3
)

//Generated of token,length 32
func GenerateToken() string {
	bytes := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, bytes)
	if err != nil {
		log.Fatal("Generated token fail", err)
	}
	token := base64.StdEncoding.EncodeToString(bytes)
	return token
}

func Home(c *gin.Context) {

	var posts []models.Posts
	models.DB.Raw("SELECT * FROM posts ORDER BY id DESC limit ?", sp).Scan(&posts)
	if len(posts) == 0 {
		c.HTML(200, "compose.html", "")
		return
	}

	var sumPage []int
	models.DB.Raw("SELECT count(*) From posts").Scan(&sumPage)

	c.HTML(200, "home.html", gin.H{
		"posts":   posts,
		"sumPage": sumPage[0],
	})
	// c.String(200, "%s", Posts[0].Title)
}

func ComposeGet(c *gin.Context) {
	id := c.Query("id")
	var posts []models.Posts
	if len(id) > 0 {
		models.DB.Raw("SELECT * FROM posts WHERE id=?", id).Scan(&posts)
	}
	c.HTML(200, "compose.html", gin.H{
		"posts": posts,
	})

}
func LoginGet(c *gin.Context) {
	token := GenerateToken()
	// Expires := time.Now().Add(24 * time.Hour)
	// MaxAge := 86400
	var userCookie http.Cookie
	userCookie.Name = "_xsrf"
	userCookie.Value = token
	// userCookie.MaxAge = MaxAge
	http.SetCookie(c.Writer, &userCookie)
	c.HTML(200, "login.html", "")
}
func LoginPost(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	if len(email) > 0 && len(password) > 0 {
		var user []models.Users
		models.DB.Raw("SELECT * FROM users WHERE email=? AND password=?", email, password).Scan(&user)
		if len(user) > 0 {
			value := base64.StdEncoding.EncodeToString([]byte(strconv.FormatUint(user[0].Uid, 10)))
			// MaxAge := time.Now().Add(24 * time.Hour).Second()
			MaxAge := 86400
			var userCookie http.Cookie
			userCookie.Name = "user"
			userCookie.Value = value
			userCookie.MaxAge = MaxAge
			http.SetCookie(c.Writer, &userCookie)
			c.String(200, "ok")
		} else {
			c.String(200, "false")
		}
	} else {
		c.String(200, "false")
	}
}

func LogoutGet(c *gin.Context) {
	var userCookie http.Cookie
	userCookie.Name = "user"
	userCookie.MaxAge = -1
	http.SetCookie(c.Writer, &userCookie)
	c.Redirect(http.StatusMovedPermanently, "/")

}

func AboutGet(c *gin.Context) {
	c.HTML(http.StatusOK, "about.html", "")
}

func ArchiveGet(c *gin.Context) {
	var posts []models.Posts
	models.DB.Raw("SELECT * FROM posts").Scan(&posts)
	c.HTML(http.StatusOK, "archive.html", gin.H{
		"posts": posts,
	})
}

// func ComposePost(c *gin.Context) {
// 	id := c.Query("id")
// 	title := c.PostForm("title")
// 	content := c.PostForm("content")

// 	if title && content {
// 		if id {
// 			var article models.Posts
// 			models.DB.Raw("SELECT * FROM posts WHERE id=?", id).Scan(&article)
// 			if !article {
// 				c.HTML(404, "404.html", "")
// 				return
// 			}
// 			slug := article.slug
// 			models.DB.Raw(
// 				"UPDATE posts SET title = ? content = ? WHERE id =?", title, content, id)
// 		} else {
// 			var max string
// 			models.DB.Raw("SELECT MAX(*) From posts").Scan(&max)
// 			date := time.Now().Format("20060102")
// 			slug := date + max
// 			models.DB.Raw(
// 				"INSERT INTO posts (title, slug, content, published) VALUES (?,?,?,UTC_TIMESTAMP())",
// 				title, slug, content)
// 		}
// 		c.Redirect(304, "home.html")
// 	} else {
// 		c.String(200, "Please enter a valid title and content")
// 	}
// }
