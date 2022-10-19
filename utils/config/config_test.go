package config

import (
	"testing"

	"github.com/spf13/viper"
)

func Test_ConfigInit(t *testing.T) {
	ConfigInit()
	t.Log(viper.Get("server.runMode"))
	t.Log(viper.Get("log.output"))
	t.Log(viper.Get("mysql.host"))
}
