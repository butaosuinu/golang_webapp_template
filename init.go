package main

import (
	"./controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db gorm.DB

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*")

	router.GET("/", controller.IndexView)
	// router.GET("/login", controller.LoginView)
	// router.POST("/login", LoginUser)
	// router.GET("/user/:id", ViewUserPage)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	router.Static("/public", "public")
	router.Run(":8080")
}

func init() {
	db, _ = gorm.Open("mysql", "root@/gorm?charset=utf8&parseTime=True")
}
