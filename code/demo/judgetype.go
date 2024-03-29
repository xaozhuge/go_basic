package main

import (
	"code/P"
)

func typeInt(){
	P.PE("typeInt")
	temp_var := 42
	P.D(temp_var)
}

func typeFloat64(){
	P.PE("typeFloat64")
	temp_var := 3.14
	P.D(temp_var)
}

func typeString(){
	P.PE("typeString")
	temp_var := "Hello, World!"
	P.D(temp_var)
}

func typeBool(){
	P.PE("typeBool")
	temp_var := true
	P.D(temp_var)
}

// 定义结构体
type Person struct {
    Name string
    Age  int
}

func typeStruct(){
	P.PE("typeStruct")
	temp_var := Person{"John", 25}
	P.D(temp_var)
}



func main() {
	typeInt()
	typeFloat64()
	typeString()
	typeBool()
	typeStruct()
}

/*
docker exec go_c go run demo/judgetype.go
*/

/*
1. 使用 reflect.TypeOf() 判断变量类型
*/