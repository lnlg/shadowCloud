package redis

import (
	"context"
	"shadowCloud/internal/global"

	"github.com/redis/go-redis/v9"
)

func New() (Rdb *redis.Client) {
	config := global.Config.Redis
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.Host + ":" + config.Port,
		Password: config.Password,
		DB:       config.Db,
	})

	if err := Rdb.Ping(context.Background()).Err(); err != nil {
		panic("redis 连接失败！")
	}
	return Rdb
}
