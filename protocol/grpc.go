package protocol

import (
	"context"
	"net"
	"time"

	"google.golang.org/grpc"

	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcenter/apps/instance"
	"github.com/infraboard/mcenter/client/rpc"
	"github.com/infraboard/mcenter/client/rpc/middleware"

	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/grpc/middleware/recovery"
	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"
)

// NewGRPCService todo
func NewGRPCService() *GRPCService {
	rc := recovery.NewInterceptor(recovery.NewZapRecoveryHandler())
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		rc.UnaryServerInterceptor(),
		middleware.GrpcAuthUnaryServerInterceptor(),
	))

	// 控制Grpc启动其他服务, 比如注册中心
	ctx, cancel := context.WithCancel(context.Background())

	return &GRPCService{
		svr: grpcServer,
		l:   zap.L().Named("server.grpc"),
		c:   conf.C(),

		ctx:    ctx,
		cancel: cancel,
		client: rpc.C(),
	}
}

// GRPCService grpc服务
type GRPCService struct {
	svr *grpc.Server
	l   logger.Logger
	c   *conf.Config

	ctx    context.Context
	cancel context.CancelFunc
	ins    *instance.Instance
	client *rpc.ClientSet
}

// Start 启动GRPC服务
func (s *GRPCService) Start() {
	// 装载所有GRPC服务
	app.LoadGrpcApp(s.svr)

	// 启动HTTP服务
	lis, err := net.Listen("tcp", s.c.App.GRPCAddr())
	if err != nil {
		s.l.Errorf("listen grpc tcp conn error, %s", err)
		return
	}

	time.AfterFunc(1*time.Second, s.registry)

	s.l.Infof("GRPC 服务监听地址: %s", s.c.App.GRPCAddr())
	if err := s.svr.Serve(lis); err != nil {
		if err == grpc.ErrServerStopped {
			s.l.Info("service is stopped")
		}

		s.l.Error("start grpc service error, %s", err.Error())
		return
	}
}

func (s *GRPCService) registry() {
	req := instance.NewRegistryRequest()
	req.Address = s.c.App.GRPCAddr()
	ins, err := s.client.Instance().RegistryInstance(s.ctx, req)
	if err != nil {
		s.l.Errorf("registry to mcenter error, %s", err)
		return
	}
	s.ins = ins

	s.l.Infof("registry instance to mcenter success")
}

// Stop 启动GRPC服务
func (s *GRPCService) Stop() error {
	// 提前 剔除注册中心的地址
	if s.ins != nil {
		req := instance.NewUnregistryRequest(s.ins.Id)
		if _, err := s.client.Instance().UnRegistryInstance(s.ctx, req); err != nil {
			s.l.Errorf("unregistry error, %s", err)
		} else {
			s.l.Info("unregistry success")
		}
	}

	s.svr.GracefulStop()

	s.cancel()
	return nil
}
