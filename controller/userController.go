package controller

import (
	"github.com/gin-gonic/gin"
	"jwtTest/jwt"
	"net/http"
	"time"
)

type User struct {
	UserID   int
	Username string
}

func Login(context *gin.Context) {
	var user User
	if err := context.BindJSON(&user); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	payload := jwt.Payload{
		Issuer:     "youngalone",
		IssuedAt:   time.Now().String(),
		Expiration: time.Now().Add(time.Hour * 24).String(),
		Audience:   user.Username,
	}
	sign, err := jwt.Sign(jwt.Header{Alg: "HS256", Typ: "JWT"}, payload)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, err)
	}
	context.JSON(http.StatusOK, sign)
}

func Doing(context *gin.Context) {
	username, _ := context.Get("Username")
	context.JSON(http.StatusOK, username)
}
