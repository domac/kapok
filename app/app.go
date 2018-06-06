package app

import (
	"github.com/codegangsta/cli"
	kapok "github.com/domac/kapok/core"
	"github.com/domac/kapok/util"
	"os"
)

//应用启动Action
func appAction(c *cli.Context) (err error) {
	//以命令行的方式启动
	target := os.Args[len(os.Args)-1]
	return kapok.CreatePlayLoad(c, target)
}

//程序主入口
func Startup() {
	flagsInit()
	app := cli.NewApp()
	app.Name = "kapok"
	app.Usage = "a simple http/https benchmark utility"
	app.Version = APP_VERSION
	app.Flags = util.GetAppFlags()
	app.Action = util.ActionWrapper(appAction)
	app.Run(os.Args)
}
