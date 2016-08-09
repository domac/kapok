package core

import (
	"fmt"
	"github.com/phillihq/kapok/util"
	"time"
)

type Stats struct {
	Url            string
	RespSize       int64
	Duration       time.Duration
	MinRequestTime time.Duration
	MaxRequestTime time.Duration
	NumRequests    int
	NumErrs        int
	Num5X          int
	Num2X          int
}

//输出统计信息
func (aggStats *Stats) PrintStats(responders int) (res string) {
	avgThreadDur := aggStats.Duration / time.Duration(responders) //need to average the aggregated duration
	reqRate := float64(aggStats.NumRequests) / avgThreadDur.Seconds()
	qpsRate := float64(aggStats.NumRequests) / (float64(responders) * avgThreadDur.Seconds())
	avgReqTime := aggStats.Duration / time.Duration(aggStats.NumRequests)
	bytesRate := float64(aggStats.RespSize) / avgThreadDur.Seconds()
	res += fmt.Sprintf("%v requests in %v, %v read\n", aggStats.NumRequests, avgThreadDur, util.ByteSize{float64(aggStats.RespSize)})
	res += fmt.Sprintf("Requests/second:\t%.2f\nTransfer/second:\t%v\nAvg Request Time:\t%v\n", reqRate, util.ByteSize{bytesRate}, avgReqTime)
	res += fmt.Sprintf("Query Per Second:\t%v\n", qpsRate)
	res += fmt.Sprintf("Fastest Request:\t%v\n", aggStats.MinRequestTime)
	res += fmt.Sprintf("Slowest Request:\t%v\n", aggStats.MaxRequestTime)
	res += fmt.Sprintf("Number of Errors:\t%v\n", aggStats.NumErrs)
	res += fmt.Sprintf("Number of 2XX:\t%v\n", aggStats.Num2X)
	res += fmt.Sprintf("Number of 5XX:\t%v\n", aggStats.Num5X)
	return
}
