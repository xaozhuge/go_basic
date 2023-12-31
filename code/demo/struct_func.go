package main 

import(
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

//接收器变量 c
//接收器类型 Cat
func (c Cat) changeName1() {
	c.Name = "酷黑猫"
	fmt.Println("changeName1 中修改后：", c)
}

//接收器变量 c
//接收器类型 *Cat---指针
func (c *Cat) changeName2() {
	c.Name = "大花猫"
	fmt.Println("changeName2 中修改后：", c)
}

func main() {
	var fan = Cat{Name: "小番茄", Age: 20}
	fmt.Println("changeName1 之前：", fan)

	fan.changeName1()
	fmt.Println("changeName1 之后：", fan)

	fmt.Println(".....................")
	fmt.Println("changeName2 之前：", fan)
	
	fan.changeName2()
	fmt.Println("changeName2 之后：", fan)

	fmt.Println(".....................")
	fan1 := new(Cat)
	fan1.changeName2()
}

/*
docker exec go_c go run demo/struct_func.go
*/

/*

1. 如果使用值接收者，方法内部对接收者的修改不会影响调用者的原始对象。
因为值接收者传递的是结构体的副本，对副本的修改不会影响原始结构体。
2. 如果使用指针接收者，方法内部对接收者的修改将直接影响调用者的原始对象。
指针接收者传递的是指向结构体的指针，通过指针可以修改原始结构体。
3. 使用值接收者会在方法调用时复制整个结构体，可能会引起性能损耗，尤其是结构体较大时。
4. 当你有一个指向结构体的指针时，可以使用该指针直接调用使用指针接收者声明的方法。
如果结构体方法使用指针接收者，Go 语言会在调用方法时隐式地将结构体取地址。
也就是方法里是指针，实例化调用时就算不是指针，也会自动将结构体取地址
5. 当一个方法的接收者使用了值接收者（不是指针接收者），并且你调用这个方法时传递的是该类型的指针，
Go 语言也会进行隐式的取值操作。



 */

