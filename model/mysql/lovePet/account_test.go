package lovepet

import (
	"context"
	"sync"
	"testing"
	"web-server/model/mysql"

	"github.com/spf13/viper"
)

var once sync.Once

func oneTestSoup(t *testing.T) {
	once.Do(func() {
		viper.SetConfigFile("../../../conf/mysql.yaml")
		if err := viper.ReadInConfig(); err != nil {
			t.Error("Read Config failed")
		}
		if err := mysql.InitDB(); err != nil {
			t.Error("get mysql instance failed")
		}
	})
}
func Test_GetAccountInfoByID(t *testing.T) {
	oneTestSoup(t)
	ctx := context.Background()
	result, err := GetAccountInfoByID(ctx, "1")
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
