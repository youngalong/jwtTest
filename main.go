package main

import (
	"github.com/gin-gonic/gin"
	"jwtTest/controller"
	"jwtTest/middleware"
)

func main() {
	r := gin.Default()
	r.POST("/login", controller.Login)
	r.POST("/Doing", middleware.Check, controller.Doing)
	r.Run(":8080")
}
