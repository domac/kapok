package routes

import (
	"github.com/codegangsta/cli"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine, cli *cli.Context) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "kapok verson "+cli.App.Version)
	})
}
