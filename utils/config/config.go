package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ConfigInit 解析文件，注册配置
func ConfigInit() {
	configFile := []string{"app", "mysql", "redis"}
	// 设置配置文件路径
	viper.AddConfigPath("conf/")
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	for _, name := range configFile {
		viper.SetConfigName(name)
		if err := viper.MergeInConfig(); err != nil {
			fmt.Printf("读取配置文件失败: %v\n", err)
		}
	}
	watchConfig()
}

// 监听配置文件是否改变,用于热更新
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("配置文件修改更新: %s\n", e.Name)
	})
}
