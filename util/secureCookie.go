package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const secret = "a94a8fe5ccb19ba61c4c0873d391e987982fbbd3"

func SetSecureCookie(w http.ResponseWriter, cookie *http.Cookie) {

	cookie.Value = createSignedValue(cookie.Name, cookie.Value)
	http.SetCookie(w, cookie)

}

func createSignedValue(name, value string) string {
	clock := time.Now().Unix()
	nowStamp := strconv.FormatInt(clock, 10)
	valueBase64 := base64.StdEncoding.EncodeToString([]byte(value))
	valueHash := createHmacHash(valueBase64)

	log.Println("cookie name", name)
	log.Println("cookie value", value)
	log.Println("cookie base64 value", valueBase64)
	log.Println("cookie base64 and hash256 value", valueHash)
	return nowStamp + "|" + name + "|" + valueBase64 + "|" + valueHash
}

func createHmacHash(s string) string {
	h := hmac.New(sha256.New, []byte(secret))
	_, err := h.Write([]byte(s))
	if err != nil {
		return ""
	}
	//v := h.Sum(nil)
	value := fmt.Sprintf("%x", h.Sum(nil))

	//16进制，需要转换为10进制
	//value := ""
	//for i := 0; i < len(v); i++ {
	//	value += fmt.Sprintf("%x", v[i])
	//}
	return value
}

func GetSecureCookie(r *http.Request, name string) (string, error) {

	rawCookie, err := r.Cookie(name)

	if err != nil {
		return "", err
	}

	strlist := strings.Split(rawCookie.Value, "|")
	//nowStamp := strlist[0]
	bs := strlist[2]

	//valueHash := strlist[3]
	//16进制，需要转换为10进制
	bsd, e := base64.StdEncoding.DecodeString(string(bs))

	if e != nil {
		return "", e
	}
	return string(bsd), nil

}
