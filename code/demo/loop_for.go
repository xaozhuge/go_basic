package main

import "fmt"

func basicFor(){
    fmt.Println("basicFor()--------------")
    for i := 0; i < 5; i++ {
        fmt.Println(i)
    }
    fmt.Println()
}

func infiniteLoop(){
    fmt.Println("infiniteLoop()--------------")
    j := 0
    for {
        fmt.Println("Infinite Loop")
        j++
        if j == 4 {
            break
        }
    }
    fmt.Println()
}

func imitateWhile(){
    fmt.Println("imitateWhile()--------------")
    m := 0
    for m < 5 {
        fmt.Println(m)
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