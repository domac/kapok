# kapok

a simple http/https benchmark utility


### 测试说明

```
$ go run main.go -h

NAME:
   kapok - a simple http/https benchmark utility

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.0.3

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --web		start the application in web
   --port "9090"	the port for web application [$KAPOK_PORT]
   --debug		open the debug mode
   -c "10"		Number of concurrent connections to use
   -d "10"		Duration of test in seconds
   --help, -h		show help
   --version, -v	print the version
```

### 安装相关依赖

```
go get github.com/gin-gonic/gin

go get github.com/codegangsta/cli

go get github.com/coreos/go-etcd
```