package model

import (
	"log"
	"web-server/model/mysql"
	"web-server/model/redis"
)

// ModelInit 数据层注册
func ModelInit() {
	// 注册mysql
	if err := mysql.InitDB(); err != nil {
		log.Fatalf("Init Mysql Fail! err = %+v", err)
	}

	//注册redis
	if err := redis.InitRedis(); err != nil {
		log.Fatalf("Init redis Fail! err = %+v", err)

	}
}
