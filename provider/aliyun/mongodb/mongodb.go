package mongodb

import (
	dds "github.com/alibabacloud-go/dds-20151201/v3/client"
	"github.com/infraboard/mcube/logger"
)

type MongoDBOperator struct {
	client *dds.Client
	log    logger.Logger
}
