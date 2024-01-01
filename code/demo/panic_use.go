package main

import "code/P"
// import ""

func arrayPanic(index int) int {
	arr := [3]int{2, 3, 4}
	return arr[index]
}


func arrayCatchPanic(index int) (ret int) {

	defer func() {
		if r := recover(); r != nil {
			P.D("Some error happened!", r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

func main() {
	// 数组越界
	// P.D(arrayPanic(5))

	// 捕获panic
	P.D(arrayCatchPanic(5))
}

/*
docker exec go_c go run demo/panic_use.go
*/

/*
1. 使用 defer 定义了异常处理的函数，在协程退出前，会执行完 defer 挂载的任务。
因此如果触发了 panic，控制权就交给了 defer。
2. 在 defer 的处理逻辑中，使用 recover，使程序恢复正常，并且将返回值设置为 -1，
在这里也可以不处理返回值，如果不处理返回值，返回值将被置为默认值 0。 
3. 返回值定义了变量，就不需要再加 return
 */