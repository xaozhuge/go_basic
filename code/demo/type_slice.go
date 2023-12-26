package main

import (
	// "fmt"
	"code/P"
)

func createSlice(){
	P.PE("createSlice")
	mySlice := []int{1, 2, 3, 4, 5}
	P.D(mySlice)

	// 切片（截取部分元素）
	subset := mySlice[1:3]
	P.D(subset)
}

func selectSliceElement(){
	P.PE("selectSliceElement")
	mySlice := []int{1, 2, 3, 4, 5}
	P.D(mySlice)
	// 获取元素
	element := mySlice[0]
	P.D(element)
}

func addSliceElement(){
	P.PE("addSliceElement")
	mySlice := []int{1, 2, 3, 4, 5}
	P.D(mySlice)

	// 在末尾添加元素
	mySlice = append(mySlice, 6)
	P.D(mySlice)

	// 在指定位置插入元素 
	mySlice = append(mySlice[:2], append([]int{7}, mySlice[2:]...)...)
	P.D(mySlice)
	
}

func deleteSliceElement(){
	P.PE("deleteSliceElement")
	mySlice := []int{1, 2, 3, 4, 5}
	P.D(mySlice)


	// 根据索引删除元素
	indexToRemove := 2
	mySlice = append(mySlice[:indexToRemove], mySlice[indexToRemove+1:]...)
	P.D(mySlice)

	// 删除开头元素
	mySlice = mySlice[1:]
	P.D(mySlice)

	// 删除末尾元素
	mySlice = mySlice[:len(mySlice)-1]
	P.D(mySlice)
}

func updateSliceElement(){
	P.PE("updateSliceElement")
	mySlice := []int{1, 2, 3, 4, 5}
	P.D(mySlice)
	// 根据索引修改元素
	mySlice[0] = 10
	P.D(mySlice)
	
}

func selectSlice(){
	P.PE("selectSlice")
	mySlice := []int{1, 2, 3, 4, 5}
	P.D(mySlice)

	// 获取切片长度
	length := len(mySlice)
	P.D(length)

	// 获取切片容量
	capacity := cap(mySlice)
	P.D(capacity)

	// 判断切片是否为空
	isEmpty := len(mySlice) == 0
	P.D(isEmpty)
}


func main() {
	createSlice()
	selectSliceElement()
	addSliceElement()
	deleteSliceElement()
	updateSliceElement()
	selectSlice()

}


/*
docker exec go_c go run demo/type_slice.go
*/

/*
1. append 第一个参数为slice,第二个参数为slice的参数类型
2. ...（三个点）被称为"ellipsis"，它用于展开切片、数组或参数列表。
3. append([]int{7}, mySlice[2:]...) 中，... 的作用是将切片 mySlice[2:] 中的
所有元素展开并传递给 append 函数。


*/

