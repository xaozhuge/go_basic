package main

import "code/P"
import "fmt"

// 使用 range 迭代数组或切片
func rangeSlice(){
    P.PE("rangeSlice")
    numbers := []int{1, 2, 3, 4, 5}
    P.D(numbers)

    for index, value := range numbers {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}

func rangeString(){
    P.PE("rangeString")
    // 使用 range 迭代字符串
    greeting := "Hello, Go!"
    P.D(greeting)
    for index, char := range greeting {
        fmt.Printf("Index: %d, Char: %c\n", index, char)
    }
}

func rangeMap(){
    P.PE("rangeMap")

    m2 := map[string]string{
        "Sam":   "Male",
        "Alice": "Female",
    }
    P.D(m2)

    for key, value := range m2 {
        P.D(key, value)
    }
}

func main() {
    rangeSlice()
    rangeString()
    rangeMap()

}

/*
docker exec go_c go run demo/loop_range.go
*/

/*
1. %d 用于格式化输出整数
2. %c 用于字符
3. range 循环按照Unicode码点来迭代字符串，而不是按照字节。
由于UTF-8的变长编码特性，一些字符占用的字节数可能不同。
 */