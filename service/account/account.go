package account

import (
	"fmt"
	lovepet "web-server/model/mysql/lovePet"
	"web-server/service/entity"

	"github.com/gin-gonic/gin"
)

// GetAccountInfoByAccountNumber 通过账号查询账号信息
func GetAccountInfoByAccountNumber(ctx *gin.Context, accountNumber string) (*entity.AccountInfo, error) {
	//判断账号是否正常
	if available, err := checkAccountNumberAvailable(accountNumber); !available {
		return nil, fmt.Errorf("账号不合法！" + err.Error())
	}
	//获取账号信息
	account, err := lovepet.GetAccountInfoByID(ctx, accountNumber)
	if err != nil {
		return nil, fmt.Errorf("数据库查询错误！" + err.Error())
	}
	ret := &entity.AccountInfo{
		AccountNumber: accountNumber,
		NickName:      account.NickName,
	}
	return ret, nil
}

// 判断账号是否合理
func checkAccountNumberAvailable(accountNumber string) (bool, error) {
	return true, nil
}
