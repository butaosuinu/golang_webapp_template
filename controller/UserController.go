package controller

import (
	"github.com/cbroglie/mustache"
	"github.com/gin-gonic/gin"
)

// GetLogin is Rendering user lgin page
func GetLogin(ctx *gin.Context) {
	// view := mustache.RenderFile("", gin.H{
	// 	"": "",
	// })
}

// LoginUser is
func LoginUser(ctx *gin.Context) {

}

// GetUserPage is
func GetUserPage(ctx *gin.Context) {
	userID := ctx.Param("id")
	view, err := mustache.RenderFile("view/user/Page.html.mustache", gin.H{
		"id":          userID,
		"description": "このユーザは◯◯です。<br/>ああああ",
	})
	if err != nil {
		ctx.String(500, "Internal Server Error")
	}

	ctx.Header("Content-Type", "text/html; charset=UTF-8")
	ctx.String(200, view)
}
