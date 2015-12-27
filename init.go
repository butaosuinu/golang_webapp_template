package main

import (
	"./controller"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db gorm.DB

func main() {
	router := gin.Default()

	router.GET("/", controller.GetIndex)
	router.GET("/login", controller.GetLogin)
	// router.POST("/login", LoginUser)
	router.GET("/user/:id", controller.GetUserPage)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "pong")
	})

	router.Static("/public", "./public")
	router.StaticFile("/uikit.min.js", "./bower_components/uikit/js/uikit.min.js")
	router.Static("/jquery", "./bower_components/jquery/dist")
	router.Run(":8080")
}

func init() {
	db, _ = gorm.Open("mysql", "root@/gorm?charset=utf8&parseTime=True")
}
