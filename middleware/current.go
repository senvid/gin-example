package middleware

import (
	"ginsite/models"
	"ginsite/util"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func islogin(c *gin.Context) bool {
	// cookie.Name = "user"
		// cookie.Value = "uid"
		// ck type --- > string
		ck, err := util.GetSecureCookie(c.Request, "user")
		log.Print(ck)

		if err != nil || ck == "" {
			log.Print(" no ck current...........")
			return false
		}

		uid, e := strconv.Atoi(ck)
		if e != nil {
			return false
		}

		// select id from users where id = ?
		var user models.User
		models.DB.Where("uid = ?", uid).Find(&user)
		if user.Uid != uint(uid) {
			return false
		}
		return true
}

func Current() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Set("current", islogin(c))

		c.Next()
	}
}