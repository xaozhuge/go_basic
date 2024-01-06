// calc.go
package main
import "fmt"

func add(num1, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(add(2,3))
}

/*
docker exec go_c go run demo/calc.go
*/