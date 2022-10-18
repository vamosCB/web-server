package lovepet

import (
	"context"
	"web-server/model/mysql"

	"gorm.io/gorm"
)

// Account 帐号信息表
type Account struct {
	mysql.BaseColumn
	AccountNumber string `json:"account_number" comment:"账户ID"`
	NickName      string `json:"nick_name" comment:"用户昵称"`
}

// GetAccountInfoByID 通过ID获取帐号信息
func GetAccountInfoByID(ctx context.Context, ID int) (*Account, error) {
	var result Account
	err := mysql.DB.First(&result, ID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &result, nil
}
