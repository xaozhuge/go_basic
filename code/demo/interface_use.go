package main

import "code/P"

type Person interface {
    getName() string
}

type Student struct {
    name string
    age  int
}

func (stu *Student) getName() string {
    return stu.name+"123"
}

type Worker struct {
    name   string
    gender string
}

func (w *Worker) getName() string {
    return w.name+"456"
}

// 实例转化为接口
var _ Person = (*Student)(nil)
var _ Person = (*Worker)(nil)

func selectInterface(){
    P.PE("selectInterface")
    var p Person = &Student{
        name: "Tom",
        age:  18,
    }
    var m Person = &Worker{"Alex", "男"}
    P.D(p)
    P.D(p.getName()) // Tom

    P.D(m)
    P.D(m.getName()) // Tom
}



func main() {
    selectInterface()
}




/*
docker exec go_c go run demo/interface_use.go
*/

/*
1. 注意结构体的简写和不简写
2. p作为接口无法直接取值，例如p.name，通过方法来实现
3. p 通过哪个结构体实现，使用的是哪个结构体的方法
4. 可以用以下去监测 某个类型是否实现了某个接口的所有方法
var _ Person = (*Student)(nil)
将空值 nil 转换为 *Student 类型，再转换为 Person 接口，
如果转换失败，说明 Student 并没有实现 Person 接口的所有方法。


*/