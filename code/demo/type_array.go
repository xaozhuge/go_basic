package main

import "code/P"


func createArray() {
    P.PE("createArray")
    var arr1 [5]int     // 一维
    P.D(arr1)
	var arr2 [5][5]int // 二维 
    P.D(arr2)
    var arr3 = [5]int{1, 2, 3, 4, 5}
    P.D(arr3)
	arr4 := [5]int{1, 2, 3, 4, 5}
    P.D(arr4)
	
}

func updateArray() {
    P.PE("updateArray")
	arr := [5]int{1, 2, 3, 4, 5}
	P.D(arr) 
	for i := 0; i < len(arr); i++ {
		arr[i] += 100
	}
	P.D(arr)  // [101 102 103 104 105]
		
}

func main() {
	createArray()
	updateArray()

}

/*
docker exec go_c go run demo/type_array.go
*/

/*

1. 



*/