package main

import (
	"fmt"
	"io/ioutil"
	"bufio"
	"os"

)

func main() {
	filePath := "remark"

	// 读取整个文件的内容
	content, err := ioutil.ReadFile(filePath)
	
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// 打印文件内容
	fmt.Println(string(content))


	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	// 使用 bufio.Scanner 按行读取文件内容
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	// 检查是否有错误发生
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}
}

/*
docker exec go_c go run demo/file.go
*/
