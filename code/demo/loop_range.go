package main

import "fmt"

func main() {
    // 使用 range 迭代数组或切片
    numbers := []int{1, 2, 3, 4, 5}
    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }

    fmt.Println("----------------")

    // 使用 range 迭代字符串
    greeting := "Hello, Go!"
    for index, char := range greeting {
        fmt.Printf("Index: %d, Char: %c\n", index, char)
    }
}

/*
docker exec go_c go run demo/loop_range.go
*/