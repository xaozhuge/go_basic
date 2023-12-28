package main

import "code/P"
import "fmt"


func selectRune() {
    P.PE("selectRune")
	str := "Hello, 世界!"

	// 将字符串转换为 rune 数组
	runes := []rune(str)
	P.D(runes)

	// 遍历 rune 数组并打印每个字符的 Unicode 码点
	for _, r := range runes {
		fmt.Printf("%c - %U\n", r, r)
	}
}

func main() {
	selectRune()

}

/*
docker exec go_c go run demo/type_rune.go
*/

/*

1. rune 是一种数据类型，它用于表示 Unicode 字符。
1个汉字是一个字符，一个字母是一个字符，1个表情是一个字符
每一个字符都可以用 Unicode 表示
2. rune 实际上是 int32 的别名，用于表示一个 Unicode 码点
在字符串中，每个字符可能占用一个或多个字节，字节和编码有关
而 rune 类型则用于确保能够正确表示任何 Unicode 字符，无论其占用多少个字节。
3. rune 数组 可能是指一个包含 rune 类型元素的切片或数组。
在处理 Unicode 字符时，可以使用 rune 类型来遍历字符串的每个字符。
4. fmt.Printf 中的 %c 和 %U 格式化标识符用于格式化 Unicode 字符。
%c: 用于格式化一个字符。它会打印对应的 Unicode 字符。
%U: 用于格式化一个 Unicode 码点（code point）。
它会打印对应的 Unicode 码点，并使用大写的十六进制表示。



*/