package main

import "code/P"


func selectString() {
    P.PE("selectString")
    str := "Go语言"
    for i := 0; i < len(str); i++ {
        P.D(str[i])
        P.D(string(str[i]))
    }
}

func main() {
	selectString()

}

/*
docker exec go_c go run demo/type_string.go
*/

/*

1. 字符串是由字节组成的，每个字符可能占用多个字节，而每个字节都是 uint8 类型的整数
字母占1个字节，汉字占3个字节，然后使用 string() 将这个字节转换为字符串。
str2[2] 表达式表示取字符串 str2 中索引为 2 的字节（第三个字节）。
2. 字节的底层类型是无符号的 8 位整数，即 byte 类型。
3. 字符串的索引操作返回的是字节，而不是字符。
4. 在ASCII码表中，大写字母A的ASCII码是65，小写字母a的ASCII码是97
英文字母G的编码是71，而字母o的编码是111。
5. 汉字转化为三个字节，每个字节的编码
*/