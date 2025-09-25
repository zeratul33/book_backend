// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"devinggo/modules/system/cmd"
	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = &gcmd.Command{
		Description: `默认启动所有服务`,
		Arguments: []gcmd.Argument{
			{
				Name:   "config",
				Short:  "c",
				Brief:  "config file (default config.yaml)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return All.Func(ctx, parser)
		},
	}

	Help = &gcmd.Command{
		Name:  "help",
		Brief: "查看帮助",
		Description: `
		命令提示符
		---------------------------------------------------------------------------------
		启动服务
		>> 所有服务  [go run main.go]   热编译  [gf run main.go]
		>> 初始化配置以及其他必要文件  [go run main.go unpack]
		>> 迁移数据库文件  [go run main.go migrate:xxx]
		  [go run main.go migrate:create -name NAME] 	创建一组名为 NAME 的带时间戳的 up/down 迁移文件
		  [go run main.go migrate:goto -v 版本]      	迁移到版本 V
		  [go run main.go migrate:up [-n N]]       	应用所有或 N 个 up 迁移
		  [go run main.go migrate:down [-n N]]     	应用所有或 N 个 down 迁移
		  [go run main.go migrate:force -v 版本]  	设置版本为 V 但不运行迁移（忽略脏状态）
		>> HTTP服务  [go run main.go http]
		>> 消息队列&定时任务  [go run main.go worker]
		>> 创建新模块  [go run main.go module:create -name 模块名称]
		>> 导出新模块  [go run main.go module:export -name 模块名称]
		>> 查看帮助  [go run main.go help]
		>> 查看版本  [go run main.go version]
    `,
	}

	All = &gcmd.Command{
		Name:        "all",
		Brief:       "start all server",
		Description: "this is the command entry for starting all server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			cmd.CmdInit(ctx, parser)
			g.Log().Debug(ctx, "starting all server")
			// 需要启动的服务
			var allServers = []*gcmd.Command{cmd.Http, cmd.Worker}

			for _, server := range allServers {
				var cmd = server
				utils.SafeGo(ctx, func(ctx context.Context) {
					if err := cmd.Func(ctx, parser); err != nil {
						g.Log().Panicf(ctx, "%v start fail:%v", cmd.Name, err)
					}
				})
			}
			// 信号监听
			cmd.SignalListen(ctx, cmd.SignalHandlerForOverall)
			<-cmd.ServerCloseSignal
			cmd.ServerWg.Wait()
			g.Log().Debug(ctx, "all service successfully closed ..")
			return
		},
	}
)

func init() {
	if err := Main.AddCommand(All, cmd.Http, cmd.Version, cmd.Worker, cmd.Unpack, cmd.MigrateUp, cmd.MigrateDown, cmd.MigrateGoto, cmd.MigrateCreate, cmd.MigrateForce, cmd.CreateModule, cmd.ExportModule, cmd.ImportModule, Help); err != nil {
		panic(err)
	}
}
