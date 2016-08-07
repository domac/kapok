package app

import (
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	router "github.com/phillihq/kapok/routes"
	"github.com/phillihq/kapok/util"
	"os"
)

//应用版本
const APP_VERSION = "0.0.2"

//配置文件参数
var conff = util.AddFlagString(cli.StringFlag{
	Name:   "config",
	EnvVar: "KAPOK_CONFIG",
	Value:  "config.json",
	Usage:  "the path of your config file",
})

//应用web端口
var portf = util.AddFlagString(cli.StringFlag{
	Name:   "port",
	EnvVar: "KAPOK_PORT",
	Value:  "9090",
	Usage:  "the port for web application",
})

//debug开关
var debugf = util.AddFlagBool(cli.BoolFlag{
	Name:  "debug",
	Usage: "open the debug mode",
})

//应用启动Action
func appAction(c *cli.Context) (err error) {
	configFilePath := c.String(conff.Name)
	portFlag := c.String(portf.Name)
	debugFlag := c.Bool(debugf.Name)
	if debugFlag {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	println(configFilePath)
	r := gin.New()
	router.RegisterRoutes(r, APP_VERSION)
	return r.Run(":" + portFlag)
}

//程序主入口
func Startup() {
	app := cli.NewApp()
	app.Name = "kapok"
	app.Usage = "a synchronized proxy tool"
	app.Version = APP_VERSION
	app.Flags = util.GetAppFlags()
	app.Action = util.ActionWrapper(appAction)
	app.Run(os.Args)
}
