package model

import (
	"log"
	"web-server/model/mysql"
	"web-server/model/redis"
)

// ModelRegister 数据层注册
func ModelRegister() {
	// 注册mysql
	if err := mysql.InitDB(); err != nil {
		log.Fatalf("Init Mysql Fail! err = %+v", err)
	}

	//注册redis
	if err := redis.InitRedis(); err != nil {
		log.Fatalf("Init redis Fail! err = %+v", err)

	}
}
