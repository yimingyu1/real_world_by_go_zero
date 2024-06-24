package core

import "github.com/zeromicro/go-zero/core/stores/redis"

func NewRedis(host, redisType string) *redis.Redis {
	return redis.MustNewRedis(redis.RedisConf{
		Host: host,
		Type: redisType,
	})
}
