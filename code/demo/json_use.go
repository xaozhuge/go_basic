package main

import (
	"fmt"
	"encoding/json"
	"code/P"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}

func tmp(){
	P.T()
}

func toJson(){
	P.PE("toJson")
	// 创建一个 Go 结构体对象
	person := Person{
		Name: "张三",
		Age:  30,
		City: "New York",
	}
	P.D(person)

	// 将 Go 结构体转化为 JSON 字符串
	jsonString, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	P.D(jsonString)
	P.D(string(jsonString))
}

func jsonTo(){
	P.PE("jsonTo")
	// JSON 字符串
	jsonString := `{"name":"Alice","age":25,"city":"London"}`
	P.D(jsonString)
	P.D([]byte(jsonString))

	// 创建一个空的 Go 结构体对象
	var person Person

	// 将 JSON 字符串转化为 Go 结构体
	err := json.Unmarshal([]byte(jsonString), &person)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// 打印解析后的 Go 结构体
	P.D(person)
}


func main() {
	toJson()
	jsonTo()	
}

/*
docker exec go_c go run demo/json_use.go
*/