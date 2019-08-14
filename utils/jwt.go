package utils

import (
	"strconv"
	"time"
	"github.com/dgrijalva/jwt-go"
)

var JWTSecret = []byte("!!SECRET!!")

func GenerateJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(JWTSecret)
	return t
}

var ChatJWTSecret = []byte("318a87e5-740b-4e01-a3e8-d17996b395e9")

func GenerateChatJWT(id uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = strconv.FormatInt(int64(id), 10)
	//claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	t, _ := token.SignedString(ChatJWTSecret)
	return t
}
