package client

import (
	"google.golang.org/grpc"

	"github.com/infraboard/cmdb/app/bill"
	"github.com/infraboard/cmdb/app/host"
	"github.com/infraboard/cmdb/app/rds"
	"github.com/infraboard/cmdb/app/resource"
	"github.com/infraboard/cmdb/app/secret"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// NewClient todo
func NewClient(conf *Config) (*Client, error) {
	zap.DevelopmentSetup()

	conn, err := grpc.Dial(
		conf.address,
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(conf.Authentication),
	)
	if err != nil {
		return nil, err
	}

	return &Client{
		conf: conf,
		conn: conn,
		log:  zap.L().Named("CMDB SDK"),
	}, nil
}

// Client 客户端
type Client struct {
	conf *Config
	conn *grpc.ClientConn
	log  logger.Logger
}

// Resource todo
func (c *Client) Resource() resource.ServiceClient {
	return resource.NewServiceClient(c.conn)
}

// Host todos
func (c *Client) Host() host.ServiceClient {
	return host.NewServiceClient(c.conn)
}

// Host todos
func (c *Client) Secret() secret.ServiceClient {
	return secret.NewServiceClient(c.conn)
}

// Bill service
func (c *Client) Bill() bill.ServiceClient {
	return bill.NewServiceClient(c.conn)
}

// Rds service
func (c *Client) Rds() rds.ServiceClient {
	return rds.NewServiceClient(c.conn)
}
