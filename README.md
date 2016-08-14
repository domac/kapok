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
   0.2.0

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -k                   if keep-alives are disabled                                                                                                                                                                                                                         
   -H                   the http headers sent to the target url                                                                                                                        
   -c "10"              number of concurrent connections to use                                                                                                                        
   -d "10"              duration of test in seconds                                                                                                                                    
   -t "1000"            socket/request timeout in (ms)                                                                                                                                 
   -m "GET"             http method      
   --compress           if prevents sending the "Accept-Encoding: gzip" header                                                                                                         
   --hb                 open heartbreat watcher, it need use --web first                                                                                                                                       
   --web                start the application in web                                                                                                                                   
   --port "9090"        the port for web application [$KAPOK_PORT]                                                                                                                     
   --debug              open the debug mode                                                                                                                                                                                                     
   --help, -h           show help                                                                                                                                                      
   --version, -v        print the version 
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

二、RESTful API

```
  http://localhost:9090/bench
```

```
调用参数： 

- url:              the test url

- concurrecy:       number of concurrent connections to use

- duration:         duration of test in seconds

- timeout:          socket/request timeout in (ms)

- method:           http method， default : GET

- header:           the http headers sent to the target url

- keepalive:        if keep-alives are disabled when value is "1"  default : "0"

- compress:         if prevents sending the "Accept-Encoding: gzip" header when value is "1"  default : "0"

```