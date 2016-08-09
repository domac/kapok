package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/phillihq/kapok/core"
	"strconv"
)

func Benchmark(c *gin.Context) string {
	testUrl := getStringValue(c, "url")
	concurrecy := getInValue(c, "concurrecy")
	duration := getInValue(c, "duration")
	timeout := getInValue(c, "timeout")
	method := getStringValue(c, "method")
	header := getStringValue(c, "header")

	ka := getStringValue(c, "keepalive")
	co := getStringValue(c, "compress")

	keepalive := false
	if ka == "1" {
		keepalive = true
	}
	compress := false
	if co == "1" {
		compress = true
	}

	res, err := core.Playload(testUrl, concurrecy, duration, timeout, method, header, keepalive, compress)
	if err != nil {
		return err.Error()
	}
	return res
}

func getInValue(c *gin.Context, key string) int {
	qk, ex := c.GetQuery(key)
	if !ex {
		return 0
	}
	val, err := strconv.Atoi(qk)
	if err != nil {
		return 0
	}
	return val
}

func getStringValue(c *gin.Context, key string) string {
	qk, ex := c.GetQuery(key)
	if !ex {
		return ""
	}
	return qk
}
