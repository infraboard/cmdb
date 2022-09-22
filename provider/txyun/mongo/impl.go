package mongo

import (
	mongodb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/mongodb/v20190725"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

func NewMongoOperator(client *mongodb.Client) *MongoOperator {
	return &MongoOperator{
		client: client,
		log:    zap.L().Named("tx.mongodb"),
	}
}

type MongoOperator struct {
	client *mongodb.Client
	log    logger.Logger
}
