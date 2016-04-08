package util

import (
	"io"
	"log"
	"encoding/base64"
	"crypto/rand"
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
