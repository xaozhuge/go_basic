package main

import "fmt"

func createMap() {
    // 创建映射
    person := map[string]int{
        "John": 25,
        "Alice": 30,
    }
    fmt.Println(person) // 输出: map[Alice:30 John:25]
    fmt.Println("----------")
}

func queryMap() {
    person := map[string]int{
        "John": 25,
        "Alice": 30,
    }
    // 读取映射元素
    fmt.Println("John's Age:", person["John"]) // 输出: John's Age: 25
    fmt.Println("----------")
}

func updateMap() {
    person := map[string]int{
        "John": 25,
        "Alice": 30,
    }
    // 修改映射元素
    person["John"] = 26

    fmt.Println(person) // 输出: map[Alice:30 John:26]
    fmt.Println("----------")
}



func deleteMap() {
    person := map[string]int{
        "John": 25,
        "Alice": 30,
    }
    // 删除映射元素
    delete(person, "John")
    fmt.Println(person) // 输出: map[Alice:30]
}


func main() {
	createMap()
	queryMap()
	updateMap()
	deleteMap()
}

/*
docker exec go_c go run demo/type_map.go
*/

/*
1. 映射（Map）是一种集合类型，用于存储键值对。
2. delete 属于 map的关键字，命名时需要注意
*/