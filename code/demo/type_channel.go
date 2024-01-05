package main

import "code/P"


func createChan() {
    P.PE("createChannel")
    ch := make(chan int)
    P.D(ch)
	
}



func main() {
	createChan()

}

/*
docker exec go_c go run demo/type_channel.go
*/

/*

1. channel 的操作符是 <- ，箭头的指向就是数据的流向
ch <- v    // 发送值v到Channel ch中
v := <-ch  // 从Channel ch中接收数据，并将数据赋值给v



*/