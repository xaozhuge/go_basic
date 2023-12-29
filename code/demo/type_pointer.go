package main

import "code/P"


func createPointer() {
    P.PE("createPointer")
    str := "Golang"
    P.D(str)

	var p *string = &str // p 是指向 str 的指针
	P.D(p)

	q := &str
	P.D(q)

	var m *string 
	P.D(m)

	*p = "Hello"
	P.D(*p)
	P.D(str) // Hello 修改了 p，str 的值也发生了改变

	a := *p
	P.D(a)
}

func valueAdd(num int) {
	P.PE("valueAdd")
	P.D(num)
	num += 1
	P.D(num)
}

func pointerAdd(num *int) {
	P.PE("pointerAdd")
	P.D(*num)
	P.D(num)
	*num += 1
	P.D(*num)
	P.D(num)
}



func main() {
	createPointer()
	num := 99
	P.D(num)
	valueAdd(num)
	P.D(num)
	pointerAdd(&num)
	P.D(num)
}

/*
docker exec go_c go run demo/type_pointer.go
*/

/*

1. *类型，代表类型是指针，例如 *string
*变量，代表类型是指针对应的值的类型
&变量，代表类型是指针，是变量的指针
2. 指针类型的空值是 nil，例如*string的空值是 nil
3. 指针即某个值的地址，类型定义时使用符号*，
对一个已经存在的变量，使用 & 获取该变量的地址。
4. 






*/