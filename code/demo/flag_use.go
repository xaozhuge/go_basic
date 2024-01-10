package main

import (
	"code/P"
	"flag"
)

var (
	host       string
	localPort  int
	remotePort int
)

func init() {
	flag.StringVar(&host, "h", "127.0.0.1", "remote server ip")
	flag.IntVar(&localPort, "l", 8080, "the local port")
	flag.IntVar(&remotePort, "r", 3333, "remote server port")
}


func main() {
	flag.Parse()
	P.D(host)
	P.D(localPort)
	P.D(remotePort)
}

/*
docker exec go_c go run demo/flag_use.go
docker exec go_c go run demo/flag_use.go -h 127.0.0.2 -l 8081 -r 3334
*/

/*
1. flag.StringVar(): 它用于定义一个字符串类型的命令行标志（flag）并关联一个变量。
在这里，&host 表示将命令行参数解析后的值存储到 host 变量中，
"h" 表示命令行参数的名称，"127.0.0.1" 是默认值，用于在用户未指定参数时使用，
最后的字符串是关于这个标志的说明。
2. 用户可以通过运行程序时提供类似 -h 192.168.1.1 -l 8080 -r 4444 
这样的命令行参数来定制这些值。如果用户没有提供相应的参数，那么将使用默认值。
3. -h 和 数值之间必须有空格
4. flag.Parse() 【必须调用】, 不然拿不到命令行中的值
从 arguments 中解析注册的 flag
    
  
*/