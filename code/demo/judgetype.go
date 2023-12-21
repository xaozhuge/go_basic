package main

import (
	"fmt"
	"reflect"
)

func typeInt(){
	fmt.Println("typeInt()------------")
	temp_var := 42
	fmt.Println("值为:", temp_var, "，类型为:", reflect.TypeOf(temp_var))
	fmt.Println()
}

func typeFloat64(){
	fmt.Println("typeFloat64()------------")
	temp_var := 3.14
	fmt.Println("值为:", temp_var, "，类型为:", reflect.TypeOf(temp_var))
	fmt.Println()
}

func typeString(){
	fmt.Println("typeString()------------")
	temp_var := "Hello, World!"
	fmt.Println("值为:", temp_var, "，类型为:", reflect.TypeOf(temp_var))
	fmt.Println()
}

func typeBool(){
	fmt.Println("typeBool()------------")
	temp_var := true
	fmt.Println("值为:", temp_var, "，类型为:", reflect.TypeOf(temp_var))
	fmt.Println()
}

func main() {
	typeInt()
	typeFloat64()
	typeString()
	typeBool()
}

/*
docker exec go_c go run demo/judgetype.go
*/

/*
1. 使用 reflect.TypeOf() 判断变量类型
*/