package redis

import (
	"context"

	goredis "github.com/redis/go-redis/v9"
)

func NewRedisClient(addr string) (*goredis.Client, error) {
	rdb := goredis.NewClient(&goredis.Options{
		Addr: addr,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
