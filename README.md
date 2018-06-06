# kapok

简单高效的HTTP压测工具，支持GET和POST等方法处理，支持HTTP Header自定义请求


### 使用说明

```
$ go run main.go -help

NAME:
   kapok - a simple http/https benchmark utility

USAGE:
   main [global options] command [command options] [arguments...]

VERSION:
   0.1.1

COMMANDS:
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -c "10"		number of concurrent connections to use
   -t "1000"		socket/request timeout in (ms)
   -m "GET"		http method
   --disableka		disable keep-alives
   --debug		open the debug mode
   -d "10"		duration of test in seconds
   -H 			the http headers sent to the target url
   --compress		if prevents sending the "Accept-Encoding: gzip" header
   --dataFile 		load the par which store in the file
   --help, -h		show help
   --version, -v	print the version
```

#### 运行参考

一、GET默认方式

```
 $ ./kapok -c 200 -d 60 http://192.168.100.101/test/data/send.do
```

二、POST方式

```
$ ./kapok -m POST  -c 2 -d 30 -dataFile=/apps/body.json 'http://192.168.100.101/test/data/send.do'
```

三、自定义Heaer方式

```
$ ./kapok  -c 2 -d 30 -dataFile=/apps/body.json -H  'Content-Seq:001;App-Type:phone'  'http://192.168.100.101/test/data/send.do'
```

结果输出示例：

```
Running 30s http://192.168.100.101/test/data/send.do                                                                                                                                                     
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
