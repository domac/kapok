package core

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"net/url"
	"os"
	"os/signal"
	"time"
)

//生成载荷
func CreatePlayLoad(c *cli.Context, url string) error {
	conn_num := c.Int("c")
	duration := c.Int("d")
	timeout := c.Int("t")
	method := c.String("m")
	header := c.String("H")
	ka := c.Bool("k")
	compress := c.Bool("compress")
	return Playload(url, conn_num, duration, timeout, method, header, ka, compress)
}

//载荷开启
func Playload(
	testUrl string,
	concurrecy int,
	duration int,
	timeout int,
	method string,
	header string,
	ka bool,
	co bool) error {

	//检查URL的合法性
	pUrl, err := url.Parse(testUrl)
	if err != nil {
		return err
	}

	if pUrl.Host == "" {
		return errors.New("the url is incorrect!")
	}

	statsChann := make(chan *Stats, concurrecy)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	fmt.Printf("Running %vs %v\n%v connection(s) running concurrently\n", duration, testUrl, concurrecy)

	worker := NewWorker(testUrl, concurrecy, duration,
		timeout, header, method, statsChann, ka, co)

	for i := 0; i < concurrecy; i++ {
		go worker.RunSingleNode()
	}

	responders := 0
	aggStats := Stats{MinRequestTime: time.Minute}
	for responders < concurrecy {
		select {
		case <-sigChan:
			worker.Stop()
			fmt.Printf("stop working ...\n")
		case stats := <-statsChann:
			aggStats.NumErrs += stats.NumErrs
			aggStats.NumRequests += stats.NumRequests
			aggStats.RespSize += stats.RespSize
			aggStats.Duration += stats.Duration
			aggStats.MaxRequestTime += stats.MaxRequestTime
			aggStats.MinRequestTime += stats.MinRequestTime
			responders++
		}
	}

	fmt.Printf("Finish %v concurrecy load! \n", responders)
	if aggStats.NumRequests == 0 {
		return errors.New("No statistics collected / no requests found ! ")
	}
	fmt.Println(aggStats.PrintStats(responders))
	return nil
}
