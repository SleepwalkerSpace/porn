package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Connect 连接Redis数据库 https://github.com/go-redis/redis
func Connect(addr, pswd string, db int) (*redis.Client, error) {
	cli := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pswd,
		DB:       db,
	})
	if err := cli.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return cli, nil
}
