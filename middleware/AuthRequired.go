package middleware

import (
	"github.com/gin-gonic/gin"
	"ginsite/util"
	"net/http"
	"ginsite/models"
	"strconv"
	"log"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// cookie.Name = "user"
		// cookie.Value = "uid"
		// ck type --- > string
		ck,err := util.GetSecureCookie(c.Request,"user")
		log.Print(ck)

		if err != nil || ck == "" {
			log.Print(" no ck ...........")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.AbortWithStatus(401)
			return
		}

		uid,e := strconv.Atoi(ck)
		if e != nil {
			c.Abort()
			log.Print(" no uid ...........")
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		// select id from users where id = ?
		var user models.User
		models.DB.Where("uid = ?", uid).Find(&user)
		if user.Uid != uint(uid) {
			c.Abort()
			log.Print(" no correct ck ...........")
			c.Redirect(http.StatusMovedPermanently, "/login")
			return
		}

		c.Next()
	}
}