package main

import (
	"fmt"
	"time"
)

func main() {
	// 设置时区为东八区
	location := time.FixedZone("CST", 8*60*60)
	// 获取当前时间
	now := time.Now().In(location)
	// 输出东八区时间
	fmt.Println("东八区时间:", now.Format("2006-01-02 15:04:05"))

	// 通过 time.LoadLocation 加载亚洲上海时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("无法加载时区:", err)
		return
	}

	// 获取当前时间在上海时区的时间
	now_time := time.Now().In(location)

	// 输出上海时区时间
	fmt.Println("上海时区时间:", now_time.Format("2006-01-02 15:04:05"))


}

/*
docker exec go_c go run demo/standard_time.go
*/




	
