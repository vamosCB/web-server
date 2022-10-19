package main

import (
	"fmt"
	"web-server/model"
	"web-server/utils/config"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 初始化配置
	config.ConfigInit()
	// 初始化数据库
	model.ModelInit()
	// // 访问/version的返回值会随配置文件的变化而变化
	// r.GET("/version", func(c *gin.Context) {
	// 	c.String(http.StatusOK, Conf.Version)
	// })
	if err := r.Run(fmt.Sprintf(":%d", 8080)); err != nil {
		panic(err)
	}
}
