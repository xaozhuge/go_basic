package main

import "fmt"

// 定义结构体
type Person struct {
    Name string
    Age  int
}

func createStruct() {
    // 创建结构体实例
    person1 := Person{"John", 25}
    person2 := Person{Name: "Alice", Age: 30}

    fmt.Println(person1) // 输出: {John 25}
    fmt.Println(person2) // 输出: {Alice 30}
    fmt.Println("-----------------")
}

func queryStruct() {
    person := Person{"John", 25}
    // 读取结构体字段
    fmt.Println("Name:", person.Name)
    fmt.Println("Age:", person.Age)
    fmt.Println("-----------------")
}


func updateStruct() {
    person := Person{"John", 25}
    // 修改结构体字段
    person.Name = "Alice"
    person.Age = 30
    fmt.Println(person) // 输出: {Alice 30}
    fmt.Println("-----------------")
}

//结构体的字段是固定的，不能直接删除字段。
//如果需要删除字段，可以通过创建一个新的结构体实例，只包含需要的字段，来模拟删除操作。
func deleteStruct() {
    person1 := Person{"John", 25}
    // 模拟删除 Name 字段
    person2 := Person{Age: person1.Age}

    fmt.Println(person2)
}

func main() {
	createStruct()
	queryStruct()
	updateStruct()
	deleteStruct()
}

/*
docker exec go_c go run demo/type_struct.go
*/

/*
1. 结构体是一种用户自定义的数据类型，用于表示一组相关的字段。
*/