package controllers

import (
	. "ginsite/models"
	"ginsite/util"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func SetXsrf(c *gin.Context) string {
	xsrf, err := c.Cookie("_xsrf")
	if err != nil || xsrf == "" {
		token := util.GenerateToken()
		userCookie := http.Cookie{
			Name:     "_xsrf",
			Value:    token,
			MaxAge:   86400,
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, &userCookie)
	}
	log.Println("_xsrf: ", xsrf)
	return xsrf
}
func LoginHandler(c *gin.Context) {
	//Email := c.PostForm("email")
	//Password:=c.PostForm("password")
	//if Email=="" || Password=="" {
	//	c.HTML(http.StatusOK,"login.html",gin.H{})
	//	return
	//}
	_ = SetXsrf(c)
	c.HTML(http.StatusOK, "login.html", gin.H{})

}

func LoginPostHandler(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")
	xsrf, err := c.Cookie("_xsrf")
	if err != nil || xsrf == "" {
		c.String(http.StatusForbidden, "false")
		return
	}
	if email != "" && password != "" {
		var user User
		DB.Where("email = ? and password = ?", email, password).Find(&user)

		if user.Uid == 0 {
			c.String(http.StatusUnauthorized, "false")
			return
		}

		//value := base64.StdEncoding.EncodeToString([]byte(strconv.Itoa(int(user.Uid))))
		value := strconv.Itoa(int(user.Uid))
		userCookie := http.Cookie{
			Name:     "user",
			Value:    value,
			MaxAge:   86400,
			HttpOnly: true,
		}
		util.SetSecureCookie(c.Writer, &userCookie)
		//c.String(200, "ok")
		c.Redirect(301, "/")
	} else {
		c.String(http.StatusUnauthorized, "false")
	}

}

func LogoutHandler(c *gin.Context) {
	var userCookie http.Cookie
	userCookie.Name = "user"
	userCookie.MaxAge = -1
	http.SetCookie(c.Writer, &userCookie)
	c.Redirect(http.StatusMovedPermanently, "/")
}
