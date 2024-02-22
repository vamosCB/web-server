package controller

import (
	"fmt"
	"net/http"
	"web-server/service/account"
	"web-server/service/entity"

	"github.com/gin-gonic/gin"
)

// GetAccountInfoByAccountNumber 通过账号查询账户信息
func GetAccountInfoByAccountNumber(ctx *gin.Context) {
	param := &entity.AccountNumberParam{}
	if err := ctx.ShouldBind(param); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("参数错误"))
	}
	_, err := account.GetAccountInfoByAccountNumber(ctx, param.AccountNumber)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, fmt.Errorf("获取账户信息错误"))
	}
	//ctx.JSON(http.StatusOK, ret)
}
