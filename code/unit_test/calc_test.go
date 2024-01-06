// calc_test.go
package main

import "testing"

func TestAdd(t *testing.T) {
	if ans := add(1, 2); ans != 3 {
		t.Error("add(1, 2) should be equal to 3")
	}
}

/*
docker exec go_c sh -c "cd unit_test;go test"
docker exec go_c sh -c "cd unit_test;go test -v"
*/


/*
1. 运行 go test，将自动运行当前 package 下的所有测试用例，
如果需要查看详细的信息，可以添加-v参数。


 */