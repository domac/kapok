package main

import (
	"github.com/phillihq/kapok/app"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	app.Main()
}
