# kapok

a simple http/https benchmark utility


### 测试说明

```
$ go run main.go -h

NAME:
   kapok - a simple http/https benchmark utility

USAGE:
   ./kapok [global options] command [command options] [arguments...]

VERSION:
   0.0.3

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -c "10"		number of concurrent connections to use
   -d "10"		duration of test in seconds
   -m "GET"		http method
   -t "1000"		socket/request timeout in (ms)
   -H 			the http headers sent to the target url
   -k			if keep-alives are disabled
   --debug		open the debug mode
   --compress		if prevents sending the "Accept-Encoding: gzip" header
   --web		start the application in web
   --port "9090"	the port for web application [$KAPOK_PORT]
   --help, -h		show help
   --version, -v	print the version
```