package main

import (
	"time"
	"fmt"
	"code/P"
)

var ch = make(chan string, 10) // 创建大小为 10 的缓冲信道

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch <- url // 将 url 发送给信道
}

func main() {
	startTime := time.Now()
	for i := 0; i < 3; i++ {
		go download("a.com/" + string(i+'0'))
	}
	for i := 0; i < 3; i++ {
		msg := <-ch // 等待信道返回消息。
		fmt.Println("finish", msg)
	}
	fmt.Println("Done!")
	elapsedTime := time.Since(startTime)
	P.D(elapsedTime)
}


/*
docker exec go_c go run demo/channel_use.go
*/

/*
1. 





 
 */