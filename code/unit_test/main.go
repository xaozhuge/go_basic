// main.go
package main

import "fmt"

func main() {
	fmt.Println(add(3,5))
}

/*
docker exec go_c go run unit_test/main.go unit_test/calc.go
或者
docker exec go_c sh -c "cd unit_test;go run ."
*/


/*
1. Go 语言也有 Public 和 Private 的概念，粒度是包。
如果类型/接口/方法/函数/字段的首字母大写，则是 Public 的，对其他 package 可见，
如果首字母小写，则是 Private 的，对其他 package 不可见。


*/