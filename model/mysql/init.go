package mysql

import (
	"fmt"

	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DB 数据库实例
var DB *gorm.DB

// BaseColumn 基础字段
type BaseColumn struct {
	ID        int       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at" comment:"创建时间"`
	UpdatedAt time.Time `json:"updated_at" comment:"更新时间"`
}

// InitDB 初始化mysql数据库
func InitDB() error {
	var err error
	DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("mysql.username"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.database"),
		viper.GetString("mysql.charset")))

	if err != nil {
		return fmt.Errorf("models.Setup err: %v", err)
	}
	DB.SingularTable(true)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
	return nil
}
