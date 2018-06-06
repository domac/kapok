package core

import (
	"errors"
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/domac/kapok/util"
	"io/ioutil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//生成载荷
func CreatePlayLoad(c *cli.Context, url string) error {
	conn_num := c.Int("c")
	duration := c.Int("d")
	timeout := c.Int("t")
	method := c.String("m")
	header := c.String("H")
	disableka := c.Bool("disableka")
	compress := c.Bool("compress")

	dataFile := c.String("dataFile")

	res, err := Playload(url, conn_num, duration, timeout, method, header, disableka, compress, dataFile)
	if err != nil {
		return err
	}
	fmt.Println(res)
	return nil
}

//载荷开启
func Playload(
	testUrl string,
	concurrecy int,
	duration int,
	timeout int,
	method string,
	header string,
	disableka bool,
	co bool, dataFile string) (string, error) {

	//检查URL的合法性
	pUrl, err := url.Parse(testUrl)
	if err != nil {
		return "", err
	}

	if pUrl.Host == "" {
		return "", errors.New("the url is incorrect!")
	}

	statsChann := make(chan *Stats, concurrecy)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGSTOP)

	fmt.Printf("Running %vs %v\n%v connection(s) running concurrently\n", duration, testUrl, concurrecy)

	//body reader
	var data []byte
	if dataFile != "" && util.CheckDataFileExist(dataFile) == nil {
		f, _ := os.Open(dataFile)
		defer f.Close()
		data, _ = ioutil.ReadAll(f)
	}

	worker := NewWorker(testUrl, concurrecy, duration,
		timeout, header, method, statsChann, disableka, co, data)

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
			aggStats.Num2X += stats.Num2X
			aggStats.Num5X += stats.Num5X
			responders++
		}
	}

	fmt.Printf("Finish %v concurrecy load! \n\n", responders)
	if aggStats.NumRequests == 0 {
		return "", errors.New("No statistics collected / no requests found ! ")
	}
	return aggStats.PrintStats(responders), nil
}
