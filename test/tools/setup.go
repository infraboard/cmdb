package tools

import (
	"os"

	"github.com/infraboard/cmdb/conf"
	"github.com/infraboard/mcube/app"
	"github.com/infraboard/mcube/logger/zap"

	// 注册所有服务
	_ "github.com/infraboard/cmdb/apps"
)

func AccessToken() string {
	return os.Getenv("MCENTER_ACCESS_TOKEN")
}

func DevelopmentSetup() {

	// 初始化日志实例
	zap.DevelopmentSetup()

	// 初始化配置, 提前配置好/etc/unit_test.env
	err := conf.LoadConfigFromEnv()
	if err != nil {
		panic(err)
	}

	// 初始化全局app
	if err := app.InitAllApp(); err != nil {
		panic(err)
	}
}
