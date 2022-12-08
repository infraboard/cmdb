package rpc

import (
	"context"
	"fmt"
	"time"

	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcenter/client/rpc/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/infraboard/cmdb/apps/bill"
	"github.com/infraboard/cmdb/apps/host"
	"github.com/infraboard/cmdb/apps/rds"
	"github.com/infraboard/cmdb/apps/resource"
	"github.com/infraboard/cmdb/apps/secret"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// NewClient todo
func NewClientSet(conf *rpc.Config) (*ClientSet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 连接到服务
	conn, err := grpc.DialContext(
		ctx,
		fmt.Sprintf("%s://%s", resolver.Scheme, "cmdb"),
		grpc.WithPerRPCCredentials(rpc.NewAuthentication(conf.ClientID, conf.ClientSecret)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithBlock(),
	)

	if err != nil {
		return nil, err
	}

	return &ClientSet{
		conf: conf,
		conn: conn,
		log:  zap.L().Named("sdk.cmdb"),
	}, nil
}

// Client 客户端
type ClientSet struct {
	conf *rpc.Config
	conn *grpc.ClientConn
	log  logger.Logger
}

// Resource todo
func (c *ClientSet) Resource() resource.RPCClient {
	return resource.NewRPCClient(c.conn)
}

// Host todos
func (c *ClientSet) Host() host.ServiceClient {
	return host.NewServiceClient(c.conn)
}

// Host todos
func (c *ClientSet) Secret() secret.RPCClient {
	return secret.NewRPCClient(c.conn)
}

// Bill service
func (c *ClientSet) Bill() bill.ServiceClient {
	return bill.NewServiceClient(c.conn)
}

// Rds service
func (c *ClientSet) Rds() rds.ServiceClient {
	return rds.NewServiceClient(c.conn)
}
