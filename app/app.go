package app

import (
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"github.com/phillihq/kapok/util"
	"os"
)

//配置文件参数
var conff = util.AddFlagString(cli.StringFlag{
	Name:   "config",
	EnvVar: "KAPOK_CONFIG",
	Value:  "config.json",
	Usage:  "the path of your config file",
})

//应用启动Action
func appAction(c *cli.Context) (err error) {
	configFilePath := c.String(conff.Name)
	println(configFilePath)
	r := gin.New()
	return r.Run(":9000")
}

//程序主入口
func Main() {
	app := cli.NewApp()
	app.Name = "kapok"
	app.Usage = "a synchronized proxy tool"
	app.Version = "0.0.1"
	app.Flags = util.GetAppFlags()
	app.Action = util.ActionWrapper(appAction)
	app.Run(os.Args)
}
