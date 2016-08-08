package app

import (
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	router "github.com/phillihq/kapok/routes"
	"github.com/phillihq/kapok/util"
	"os"
)

//应用启动Action
func appAction(c *cli.Context) (err error) {
	portFlag := c.String("port")
	debugFlag := c.Bool("debug")
	webFlag := c.Bool("web")
	if webFlag == false {
		//以命令行的方式启动
		println("========== cli")
		return nil
	} else {
		//以web的方式启动
		if debugFlag {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
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
	app.Usage = "a data service tool"
	app.Version = APP_VERSION
	app.Flags = util.GetAppFlags()
	app.Action = util.ActionWrapper(appAction)
	app.Run(os.Args)
}
