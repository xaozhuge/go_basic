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


*/