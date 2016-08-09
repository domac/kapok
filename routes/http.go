package routes

import (
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"github.com/phillihq/kapok/handler"
	"net/http"
)

func RegisterRoutes(r *gin.Engine, cli *cli.Context) {
	//设置静态资源
	r.Static("/static", "./static")

	//首页
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "kapok verson "+cli.App.Version)
	})

	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/static/html/index.html")
	})

	//性能测试API
	r.GET("/bench", func(c *gin.Context) {
		c.String(http.StatusOK, handler.Benchmark(c))
	})

}
