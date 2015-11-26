package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", IndexView)

	router.GET("/ping", func(ctx gin.Context) {
		ctx.String(200, "pong")
	})

	router.Run(":8080")
}
