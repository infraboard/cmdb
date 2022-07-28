package redis

import (
	redis "github.com/alibabacloud-go/r-kvstore-20150101/v2/client"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewRedisOperator(client *redis.Client) *RedisOperator {
	return &RedisOperator{
		client: client,
		log:    zap.L().Named("ALI Redis"),
	}
}

type RedisOperator struct {
	client *redis.Client
	log    logger.Logger
}
