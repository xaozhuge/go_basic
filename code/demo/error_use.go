package main

import (
	"code/P"
	"os"
	"errors"
)

func selectError(){
	P.PE("selectError")
	_, err := os.Open("filename.txt")
	if err != nil {
		P.D(err)
	}
}

func hello(name string) error {

	if len(name) == 0 {
		return errors.New("error: name is null")
	}

	P.D("Hello,", name)
	return nil

}

func selectSelfError(){
	P.PE("selectSelfError")
	if err := hello(""); err != nil {
		P.D("fail",err)
	}
}

func main() {
	selectError()
	selectSelfError()
}






/*
docker exec go_c go run demo/error_use.go
*/

/*
1. if 不使用 括号
2. if 里，可以写 变量赋值

*/