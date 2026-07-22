package utils

import (
	schedulerv1 "github.com/parbhat-cpp/koko-go/proto/gen/scheduler/v1"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/types/known/durationpb"
)

func ToSchedulerRedisOptions(cfg *redis.Options) *schedulerv1.RedisOptions {
	return &schedulerv1.RedisOptions{
		Addr:         cfg.Addr,
		Username:     cfg.Username,
		Password:     cfg.Password,
		Db:           int32(cfg.DB),
		MaxRetries:   int32(cfg.MaxRetries),
		DialTimeout:  durationpb.New(cfg.DialTimeout),
		ReadTimeout:  durationpb.New(cfg.ReadTimeout),
		WriteTimeout: durationpb.New(cfg.WriteTimeout),
		PoolSize:     int32(cfg.PoolSize),
	}
}
