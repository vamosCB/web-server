package entity

//AccountNumberParam 账号请求参数
type AccountNumberParam struct {
	AccountNumber string `json:"account_number" binding:"required"`
}

// AccountInfo 账户信息
type AccountInfo struct {
	AccountNumber     string `json:"account_number"`      //账号
	NickName          string `json:"nick_name"`           //昵称
	PhoneEncrypt      string `json:"phone_encrypt"`       //加密手机号
	Sex               string `json:"sex"`                 //性别
	Birthday          string `json:"birthday"`            //年龄
	Desc              string `json:"desc"`                //描述
	ProvinceID        int    `json:"province_id"`         //省份ID
	ProvinceName      string `json:"province_name"`       //省份名称
	CityID            int    `json:"city_id"`             //城市ID
	CityName          string `json:"city_name"`           //城市名称
	AccountCreateTime string `json:"account_create_time"` //账户创建时间
	AccountStatus     int    `json:"account_status"`      //账号状态
}
