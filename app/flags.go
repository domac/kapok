package app

import (
	"github.com/codegangsta/cli"
	"github.com/phillihq/kapok/util"
)

//参数初始化
func flagsInit() {
	//配置文件参数
	util.AddFlagString(cli.StringFlag{
		Name:   "config",
		EnvVar: "KAPOK_CONFIG",
		Value:  "config.json",
		Usage:  "the path of your config file",
	})

	//是否以web的形式启动
	util.AddFlagBool(cli.BoolFlag{
		Name:  "web",
		Usage: "start the application in web",
	})

	//应用web端口
	util.AddFlagString(cli.StringFlag{
		Name:   "port",
		EnvVar: "KAPOK_PORT",
		Value:  "9090",
		Usage:  "the port for web application",
	})

	//debug开关
	util.AddFlagBool(cli.BoolFlag{
		Name:  "debug",
		Usage: "open the debug mode",
	})

}
