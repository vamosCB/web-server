package router

import (
	"web-server/controller"

	"github.com/gin-gonic/gin"
)

// RouterInit 加载路由
func RouterInit(g *gin.Engine) {

	_ = g.Group("/auth")

	account := g.Group("/account")
	{
		account.GET("/getAccountInfoByAccountNumber", controller.GetAccountInfoByAccountNumber)
	}
}
