package services

import (
	"context"
	"fmt"

	"github.com/lwinmgmg/user-go/env"
	"github.com/redis/go-redis/v9"
)

func GetRedisClient(rdConf *env.RedisServer) (*redis.Client, error) {
	rdClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", rdConf.Host, rdConf.Port),
		Username: rdConf.User,
		Password: rdConf.Password,
		DB:       rdConf.Db,
	})
	rdStatus := rdClient.Ping(context.Background())
	return rdClient, rdStatus.Err()
}

func FormatThirdpartyCode(val, client_id string) string {
	return fmt.Sprintf("thirdpartyCode:%v:%v", client_id, val)
}
