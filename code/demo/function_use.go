package main

// import "fmt"
import "code/P"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func addShort(num1 int, num2 int)(ans int) {
	ans = num1 + num2
	return
}

func div(num1, num2 int) (int, int) {
	return num1 / num2, num1 % num2
}
func main() {
	quo, rem := div(100, 17)
	P.D(quo, rem)     // 5 15
	P.D(add(100, 17)) // 117
	P.D(addShort(100, 17)) // 117
}

/*
docker exec go_c go run demo/function_use.go
*/