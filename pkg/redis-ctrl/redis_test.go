package redisctrl_test

import (
	"testing"
	"time"

	"github.com/go-redis/redismock/v9"
	redisctrl "github.com/lwinmgmg/user-go/pkg/redis-ctrl"
)

func TestRedisController(t *testing.T) {
	key := "ABC"
	db, mock := redismock.NewClientMock()
	mock.Regexp().ExpectSet(key, "def", time.Minute).SetVal("def")
	rdCtrl := redisctrl.NewRedisCtrl(db)
	if err := rdCtrl.SetKey(key, "def", time.Minute, time.Second); err != nil {
		t.Errorf("Error on setting key in redis : %v", err)
	}
	mock.ExpectGet(key).SetVal("def")
	val, err := rdCtrl.GetKey(key)
	if err != nil {
		t.Errorf("Error on getting key in redis : %v", err)
	}
	if val != "def" {
		t.Errorf("Expected 'def', getting : %v", val)
	}
	mock.ExpectDel(key).SetVal(0)
	if err := rdCtrl.DelKey([]string{key}); err != nil {
		t.Errorf("Error on deleting key in redis : %v", err)
	}
}
