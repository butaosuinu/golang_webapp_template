package controller

import (
	"github.com/cbroglie/mustache"
	"github.com/gin-gonic/gin"
)

// IndexView is
func IndexView(ctx *gin.Context) {
	defer func() {
		cause := recaver()
		if cause != nil {
			ctx.JSON(500, gin.H{
				"error": cause.(error).Error(),
			})
		}
	}()

	view, err := mustache.RenderFile("index.html.mustache", gin.H{
		"name":        "TEST",
		"description": "This is test.¥n The test is successful!!",
	})
	if err != nil {
		panic(err)
	}

	ctx.String(200, view)
}
