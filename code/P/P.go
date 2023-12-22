package P

import (
    "fmt"
    "os"
    "reflect"
)

func P(temp_var ...interface{}) {
	// 使用fmt.Println输出
	fmt.Println(temp_var...)
	// 然后退出程序
	os.Exit(0)
}

func T(temp_var ...interface{}) {
    for _, element := range temp_var {
        fmt.Println("值为:",element,"\n类型为:",reflect.TypeOf(element))
    }
    os.Exit(0)
}

func PE(func_name string){
    fmt.Printf("%s()------------\n", func_name)
}


func D(temp_var ...interface{}) {
    for _, element := range temp_var {
        fmt.Printf("值为: %+v", element)
        fmt.Println(" ，类型为:",reflect.TypeOf(element))
    }
    fmt.Println()
}



