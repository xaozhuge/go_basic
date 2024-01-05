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
2. Channel类型的定义,包括三种类型的定义。
可选的<-代表channel的方向。如果没有指定方向，那么Channel就是双向的，
既可以接收数据，也可以发送数据。
chan T          // 可以接收和发送类型为 T 的数据
chan<- float64  // 只可以用来发送 float64 类型的数据
<-chan int      // 只可以用来接收 int 类型的数据



*/