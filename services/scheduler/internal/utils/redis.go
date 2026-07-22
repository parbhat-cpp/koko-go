package utils

import (
	pb "github.com/parbhat-cpp/koko-go/proto/gen/scheduler/v1"
	"github.com/redis/go-redis/v9"
)

func ToRedisOptions(cfg *pb.RedisOptions) *redis.Options {
	return &redis.Options{
		Addr:         cfg.Addr,
		Username:     cfg.Username,
		Password:     cfg.Password,
		DB:           int(cfg.Db),
		MaxRetries:   int(cfg.MaxRetries),
		DialTimeout:  cfg.DialTimeout.AsDuration(),
		ReadTimeout:  cfg.ReadTimeout.AsDuration(),
		WriteTimeout: cfg.WriteTimeout.AsDuration(),
		PoolSize:     int(cfg.PoolSize),
	}
}
