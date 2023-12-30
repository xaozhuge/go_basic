package main

import "fmt"
import "code/P"

func test(){
	P.P(123)
}


// 定义一个简单的接口
type Shape interface {
    Area() float64
}

// 实现接口的结构体1
type Circle struct {
    Radius float64
}

// 实现接口的方法1
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// 函数接受任何实现了 Shape 接口的类型
func PrintArea(s Shape) {
    fmt.Printf("Area: %f\n", s.Area())
}

func main() {
    // 创建 Circle 实例
    circle := Circle{Radius: 5}

    // 使用接口调用方法
    PrintArea(circle)

    // 创建一个接口变量
    var shape Shape


    // 动态赋值为 Circle
    shape = circle
    PrintArea(shape)

}




/*
docker exec go_c go run demo/type_interface.go
*/

/*
1. interface 是一种类型，用于定义一组方法的集合。
2. 一个结构体要被认为是实现了某个接口，必须实现该接口中定义的所有方法。
也就是说，结构体需要提供接口中所有方法的具体实现。
3. 多个类型实现了接口，如何调用接口的公共方法，通过参数为接口去调用


*/