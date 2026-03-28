package redis

import(
	goredis "github.com/redis/go-redis/v9"
)

func NewRedisClient(addr string) *goredis.Client {
	rdb := goredis.NewClient(&goredis.Options{
		Addr: addr,
	})
	return rdb
}