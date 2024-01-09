package main

import (
	"code/P"
)

var (
	host       string
	localPort  int
	remotePort int
)

func init() {
	host = "127.0.0.1"
	localPort = 8080
	localPort = 3333
}


func main() {
	P.D(host)
	P.D(localPort)
	P.D(remotePort)
}

/*
docker exec go_c go run demo/init_use.go
*/

/*
1. 字符串一定注意使用双引号
2. init 函数是一个特殊的函数，它在程序启动时被自动调用。
在这里，init 函数用于初始化程序的一些配置参数
3. init 函数没有参数和返回值。
  
*/