package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 定义不同类型的变量
	var integerVariable int = 42
	var floatVariable float64 = 3.14
	var stringVariable string = "Hello, World!"
	var boolVariable bool = true

	// 使用 reflect.TypeOf() 判断变量类型
	fmt.Println(reflect.TypeOf(integerVariable)) // int
	fmt.Println(reflect.TypeOf(floatVariable))   // float64
	fmt.Println(reflect.TypeOf(stringVariable))  // string
	fmt.Println(reflect.TypeOf(boolVariable))    // bool
}
