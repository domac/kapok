package app

import (
	"github.com/codegangsta/cli"
	kapok "github.com/domac/kapok/core"
	"github.com/domac/kapok/hb"
	router "github.com/domac/kapok/routes"
	"github.com/domac/kapok/util"
	"github.com/gin-gonic/gin"
	"os"
)

//应用启动Action
func appAction(c *cli.Context) (err error) {
	portFlag := c.String("port")
	debugFlag := c.Bool("debug")
	webFlag := c.Bool("web")
	heartbeat := c.Bool("hb")
	if webFlag == false {
		//以命令行的方式启动
		target := os.Args[len(os.Args)-1]
		return kapok.CreatePlayLoad(c, target)
	} else {
		//以web的方式启动
		if debugFlag {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}

		//开启心跳
		if heartbeat {
			go hb.OpenHeartBeat("8888")
		}

		r := gin.New()
		router.RegisterRoutes(r, c)
		return r.Run(":" + portFlag)
	}
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
