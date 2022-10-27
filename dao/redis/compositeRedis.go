package redis

import "github.com/go-redis/redis/v9"

type CompositeRedis struct {
	Redis redis.Client
	DB    interface{}
}
