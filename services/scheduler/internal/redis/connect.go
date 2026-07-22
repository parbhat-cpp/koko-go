package rdb

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func Connect(rdbCfg redis.Options) (*redis.Client, error) {
	rdb := redis.NewClient(&rdbCfg)

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return rdb, nil
}
