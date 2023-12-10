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
}

/*
docker exec go_c go run demo/struct_func.go
*/

