package controller

import "github.com/gin-gonic/gin"

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
	ctx.HTML(200, "UserPage.html", gin.H{
		"id": userID,
	})
}
