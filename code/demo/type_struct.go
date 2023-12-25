package main

import "fmt"
import "code/P"

// 定义结构体
type Person struct {
    Name string
    Age  int
}

//接收器变量 p
//接收器类型 Person
func (p Person) changeName() {
    p.Name = "李四"
    fmt.Println(p.Name)
}

func createStruct() {
    P.PE("createStruct")
    person := Person{"John", 25}
    P.D(person) // 输出: {John 25}
}

func selectStructElement() {
    P.PE("selectStructElement")
    person := Person{"John", 25}
    P.D(person)
    // 读取结构体字段
    P.D("Name:", person.Name)
    P.D("Age:", person.Age)
}


func updateStructElement() {
    P.PE("updateStructElement")
    person := Person{"John", 25}
    P.D(person)

    // 修改结构体字段
    person.Name = "Alice"
    person.Age = 30
    P.D(person) // 输出: {Alice 30}
}

//结构体的字段是固定的，不能直接删除字段。
//如果需要删除字段，可以通过创建一个新的结构体实例，只包含需要的字段，来模拟删除操作。
func deleteStructElement() {
    P.PE("deleteStructElement")
    person1 := Person{"John", 25}
    P.D(person1)
    // 模拟删除 Name 字段
    person2 := Person{Age: person1.Age}
    P.D(person2)
}

func main() {
	createStruct()
	selectStructElement()
	updateStructElement()
	deleteStructElement()
}

/*
docker exec go_c go run demo/type_struct.go
*/

/*
1. 结构体是一种用户自定义的数据类型，用于表示一组相关的字段。
2. 方法的定义方式是 func (receiver Type) methodName(parameters) returnType。
3. 在这里，(p Person) 是一个接收者（receiver），它指定了这个方法属于哪种类型。
在这个特定的例子中，p 是一个接收者变量，它的类型是 Person
4. func (p Person) changeName(){} 表示 changeName 是一个方法，属于 Person 类型的对象，
其中 p 是一个变量名，表示在方法内部可以使用这个变量来引用调用该方法的 Person 实例。
5. 如果写成 func changeName(p Person) {}，那就是一个函数而不是方法。
在这种情况下，需要明确传递一个 Person 类型的实例作为参数，
而在方法的写法中，接收者 p 就相当于函数的第一个参数。
6. 方法的写法更符合面向对象的思想，因为它将操作数据和数据本身捆绑在一起。
调用方法时，可以通过 . 运算符访问实例的字段，而不需要显式地传递实例作为参数。
这种写法也使得代码更具可读性和一致性。
*/