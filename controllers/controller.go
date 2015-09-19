package controllers

import (
	"gin-example/models"
	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	r := c.Param("key")
	// c.String(http.StatusOK, r)
	c.HTML(200, "home.html", gin.H{
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
	models.Mdb.Raw("SELECT uid,email FROM users where uid=1").Scan(&Res)

	c.HTML(200, "home.html", gin.H{
		"args": Res,
	})
}
