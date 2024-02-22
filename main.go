package main

import (
	"fmt"
	"web-server/model"
	"web-server/router"
	"web-server/utils/config"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	// 初始化配置
	config.ConfigInit()
	// 设置运行模式
	gin.SetMode(viper.GetString("server.runMode"))
	// 创建gin实例
	g := gin.Default()
	// 添加授权IP 如不需要可使用Engine.SetTrustedProxies(nil)
	g.SetTrustedProxies([]string{"192.168.1.2"})

	// 初始化数据库
	model.ModelInit()
	// 加载路由
	router.RouterInit(g)
	// 启动服务
	if err := g.Run(fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt64("server.port"))); err != nil {
		panic(err)
	}
}
