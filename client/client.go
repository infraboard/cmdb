package client

import (
	"google.golang.org/grpc"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/apps/secret"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// NewClient todo
func NewClient(conf *Config) (*ClientSet, error) {
	zap.DevelopmentSetup()

	conn, err := grpc.Dial(
		conf.address,
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(conf.Authentication),
	)
	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conf: conf,
		conn: conn,
		log:  zap.L().Named("CMDB SDK"),
	}, nil
}

// Client 客户端
type ClientSet struct {
	conf *Config
	conn *grpc.ClientConn
	log  logger.Logger
}

// Resource todo
func (c *ClientSet) Resource() resource.ServiceClient {
	return resource.NewServiceClient(c.conn)
}

// Host todos
func (c *ClientSet) Host() host.ServiceClient {
	return host.NewServiceClient(c.conn)
}

// Host todos
func (c *ClientSet) Secret() secret.ServiceClient {
	return secret.NewServiceClient(c.conn)
}

// Bill service
func (c *ClientSet) Bill() bill.ServiceClient {
	return bill.NewServiceClient(c.conn)
}

// Rds service
func (c *ClientSet) Rds() rds.ServiceClient {
	return rds.NewServiceClient(c.conn)
}
