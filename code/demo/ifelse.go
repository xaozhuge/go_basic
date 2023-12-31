package main

import "code/P"


func selectIf(){
	P.PE("selectIf")
	age := 14
	if age < 18 {
		P.D("Kid")
	} else {
		P.D("Adult")
	}

	// 可以简写为：
	if age := 20; age < 18 {
		P.D("Kid")
	} else {
		P.D("Adult")
	}
	
}


func main() {
	selectIf()
}

/*
docker exec go_c go run demo/ifelse.go
*/

/*
1. if 不使用 括号
2. if 里，可以写 变量赋值

*/