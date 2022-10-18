package lovepet

import (
	"context"
	"sync"
	"testing"
	"web-server/model/mysql"
)

var once sync.Once

func oneTestSoup(t *testing.T) {
	once.Do(func() {
		if err := mysql.InitDB(); err != nil {
			t.Error("get mysql instance failed")
		}
	})
}
func Test_GetAccountInfoByID(t *testing.T) {
	oneTestSoup(t)
	if err := mysql.InitDB(); err != nil {
		t.Error(err)
	}
	ctx := context.Background()
	result, err := GetAccountInfoByID(ctx, 1)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", result)
}
