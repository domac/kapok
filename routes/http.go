package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(r *gin.Engine, verson string) {
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "kapok verson "+verson)
	})
}
