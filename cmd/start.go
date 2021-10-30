package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/infraboard/mcube/logger"
	"github.com/infraboard/mcube/logger/zap"

	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/cmdb/pkg"
	"github.com/infraboard/cmdb/pkg/host/impl"
	searcher "github.com/infraboard/cmdb/pkg/resource/impl"
	syncer "github.com/infraboard/cmdb/pkg/syncer/impl"
	"github.com/infraboard/cmdb/protocol"
)

// startCmd represents the start command
var serviceCmd = &cobra.Command{
	Use:   "start",
	Short: "Demo后端API服务",
	Long:  `Demo后端API服务`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// 初始化全局变量
		if err := loadGlobalConfig(confType); err != nil {
			return err
		}

		// 初始化全局日志配置
		if err := loadGlobalLogger(); err != nil {
			return err
		}

		// 初始化服务层 Ioc初始化
		if err := impl.Service.Config(); err != nil {
			return err
		}
		pkg.Host = impl.Service

		// Syncer Service
		if err := syncer.Service.Config(); err != nil {
			return err
		}
		pkg.Syncer = syncer.Service

		// resource Service
		if err := searcher.Service.Config(); err != nil {
			return err
		}
		pkg.Resource = searcher.Service

		// 启动服务
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)

		// 初始化服务
		svr, err := newService(conf.C())
		if err != nil {
			return err
		}

		// 等待信号处理
		go svr.waitSign(ch)
		// 启动服务
		if err := svr.start(); err != nil {
			if !strings.Contains(err.Error(), "http: Server closed") {
				return err
			}
		}
		return nil
	},
}

func newService(cnf *conf.Config) (*service, error) {
	http := protocol.NewHTTPService()
	svr := &service{
		http: http,
		log:  zap.L().Named("CLI"),
	}

	return svr, nil
}

type service struct {
	http *protocol.HTTPService
	log  logger.Logger
}

func (s *service) start() error {
	return s.http.Start()
}

// config 为全局变量, 只需要load 即可全局可用户
func loadGlobalConfig(configType string) error {
	// 配置加载
	switch configType {
	case "file":
		err := conf.LoadConfigFromToml(confFile)
		if err != nil {
			return err
		}
	case "env":
		err := conf.LoadConfigFromEnv()
		if err != nil {
			return err
		}
	case "etcd":
		return errors.New("not implemented")
	default:
		return errors.New("unknown config type")
	}
	return nil
}

// log 为全局变量, 只需要load 即可全局可用户, 依赖全局配置先初始化
func loadGlobalLogger() error {
	var (
		logInitMsg string
		level      zap.Level
	)
	lc := conf.C().Log
	lv, err := zap.NewLevel(lc.Level)
	if err != nil {
		logInitMsg = fmt.Sprintf("%s, use default level INFO", err)
		level = zap.InfoLevel
	} else {
		level = lv
		logInitMsg = fmt.Sprintf("log level: %s", lv)
	}
	zapConfig := zap.DefaultConfig()
	zapConfig.Level = level
	zapConfig.Files.RotateOnStartup = false
	switch lc.To {
	case conf.ToStdout:
		zapConfig.ToStderr = true
		zapConfig.ToFiles = false
	case conf.ToFile:
		zapConfig.Files.Name = "api.log"
		zapConfig.Files.Path = lc.PathDir
	}
	switch lc.Format {
	case conf.JSONFormat:
		zapConfig.JSON = true
	}
	if err := zap.Configure(zapConfig); err != nil {
		return err
	}
	zap.L().Named("INIT").Info(logInitMsg)
	return nil
}

func (s *service) waitSign(sign chan os.Signal) {
	for sg := range sign {
		switch v := sg.(type) {
		default:
			// 资源清理
			s.log.Infof("receive signal '%v', start graceful shutdown", v.String())
			if err := s.http.Stop(); err != nil {
				s.log.Errorf("graceful shutdown err: %s, force exit", err)
			}
			s.log.Infof("service stop complete")
			return
		}
	}
}

func init() {
	RootCmd.AddCommand(serviceCmd)
}
