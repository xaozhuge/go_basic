package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 执行 shell 命令
	// command := "ls -l;pwd;"
	command := "du -sh log"
	cmd := exec.Command("sh", "-c", command)

	// 捕获命令的输出
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}

	// 打印输出
	fmt.Printf("Output:\n%s", output)
}

/*
docker exec go_c go run demo/shell.go
*/