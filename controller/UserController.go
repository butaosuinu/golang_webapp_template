package controller

import (
	"github.com/cbroglie/mustache"
	"github.com/gin-gonic/gin"
)

// LoginView is Rendering user lgin page
func LoginView(ctx *gin.Context) {
	// view := mustache.RenderFile("", gin.H{
	// 	"": "",
	// })
}

// LoginUser is
func LoginUser(ctx *gin.Context) {

}

// ViewUserPage is
func ViewUserPage(ctx *gin.Context) {
	userID := ctx.Param("id")
	view, err := mustache.RenderFile("view/user/Page.html.mustache", gin.H{
		"id": userID,
	})
	if err != nil {
		ctx.String(500, "Internal Server Error")
	}

	ctx.Header("Content-Type", "text/html; charset=UTF-8")

	ctx.String(200, view)
}
