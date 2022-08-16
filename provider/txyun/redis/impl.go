package redis

import (
	redis "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/redis/v20180412"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewRedisOperator(client *redis.Client) *RedisOperator {
	return &RedisOperator{
		client: client,
		log:    zap.L().Named("tx.redis"),
	}
}

type RedisOperator struct {
	client *redis.Client
	log    logger.Logger
}
