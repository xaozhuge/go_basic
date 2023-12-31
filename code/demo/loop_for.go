package main

import "code/P"

func basicFor(){
    P.PE("basicFor")
    for i := 0; i < 5; i++ {
        P.D(i)
    }
}

func infiniteLoop(){
    P.PE("infiniteLoop")
    j := 0
    for {
        P.D("Infinite Loop")
        j++
        if j == 4 {
            break
        }
    }
}

func imitateWhile(){
    P.PE("imitateWhile")
    m := 0
    for m < 5 {
        P.D(m)
        m++
    }
}


func main() {
    // 基本的 for 循环
    basicFor()  
    // 无限循环  
    infiniteLoop()
    // 模拟 while 循环
    imitateWhile()
}


/*
docker exec go_c go run demo/loop_for.go
*/