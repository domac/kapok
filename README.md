# kapok

a simple http/https benchmark utility


### 测试说明

```
$ go run main.go -help

NAME:
   kapok - a simple http/https benchmark utility

USAGE:
   ./kapok [global options] command [command options] [arguments...]

VERSION:
   0.0.3

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -c "10"		    number of concurrent connections to use
   -d "10"		    duration of test in seconds
   -m "GET"		    http method
   -t "1000"	    socket/request timeout in (ms)
   -H 			    the http headers sent to the target url
   -k		    	if keep-alives are disabled
   --debug		    open the debug mode
   --compress	    if prevents sending the "Accept-Encoding: gzip" header
   --web		    start the application in web
   --port "9090"	the port for web application [$KAPOK_PORT]
   --help, -h		show help
   --version, -v	print the version
```

#### 运行参考

一、命令行运行方式

```
 $ ./pakop -c 200 -d 60 http://192.168.100.101/test/data/get.do
```

结果输出示例：

```
Running 30s http://192.168.100.101/test/data/get.do                                                                                                                                                     
20 connection(s) running concurrently                                                                                                                                                  
Finish 20 concurrecy load!                                                                                                                                                             
                                                                                                                                                                                       
441818 requests in 29.558398248s, 49.72MB read                                                                                                                                         
Requests/second (QPS):  14947.29                                                                                                                                                       
Transfer/second:        1.68MB                                                                                                                                                         
Avg Request Time:       1.338035ms                                                                                                                                                     
Fastest Request:        1m0.002972795s                                                                                                                                                 
Slowest Request:        1.234496241s                                                                                                                                                   
Number of Errors:       0      
```