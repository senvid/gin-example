package controllers

import (
	"gin-example/models"
	"github.com/gin-gonic/gin"
	// "net/http"
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

func Home(c *gin.Context) {
	var Post models.Posts
	models.DB.Raw("SELECT * FROM posts ORDER BY id DESC LIMIT ?", sp).Scan(&Post)
	// if Post == nil {
	// 	c.Redirect(304, "compose.html")
	// 	return
	// }

	// SumPage := make([]int, 0)
	var SumPage int
	models.DB.Raw("SELECT count(*) From posts").Scan(&SumPage)
	c.HTML(200, "home.html", gin.H{
		"articles": Post,
		"sumPage":  SumPage,
	})
}

func ComposeGet(c *gin.Context) {
	id := c.Query("id")
	var article models.Posts
	if id {
		models.DB.Raw("SELECT * FROM posts WHERE id=?", id).Scan(&article)
	}
	c.HTML(200, "compose.html", gin.H{
		"article": article,
	})

}
func ComposePost(c *gin.Context) {
	id := c.Query("id")
	title := c.PostForm("title")
	content := c.PostForm("content")

	if title && content {
		if id {
			var article models.Posts
			models.DB.Raw("SELECT * FROM posts WHERE id=?", id).Scan(&article)
			if !article {
				c.HTML(404, "404.html", "")
				return
			}
			slug := article.slug
			models.DB.Raw(
				"UPDATE posts SET title = ? content = ? WHERE id =?", title, content, id)
		} else {
			slug := "201591911"
			models.DB.Raw(
				"INSERT INTO posts (title, slug, content, published) VALUES (?,?,?,UTC_TIMESTAMP())",
				title, slug, content)
		}
		c.Redirect(304, "home.html")
	} else {
		c.String(200, "Please enter a valid title and content")
	}
}
