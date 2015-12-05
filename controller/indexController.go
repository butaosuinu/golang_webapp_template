package controller

import "github.com/gin-gonic/gin"

// IndexView is
func IndexView(ctx *gin.Context) {
	defer func() {
		cause := recover()
		if cause != nil {
			ctx.JSON(500, gin.H{
				"error": cause.(error).Error(),
			})
		}
	}()

	// view, err := mustache.RenderFile("view/index.html.mustache", gin.H{
	// 	"name":        "TEST",
	// 	"description": "This is test.¥n The test is successful!!",
	// })
	// if err != nil {
	// 	ctx.String(500, "Internal Server Error")
	// }

	ctx.HTML(200, "index.html", gin.H{
		"name":        "TEST",
		"description": "This is test.¥n The test is successful!!",
	})

}
