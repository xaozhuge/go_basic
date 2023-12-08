package main

import (
	"fmt"
	"code/common"//引入路径从 code 开始
	"code/P"//引入路径从 code 开始
)

func main() {
	//时区
	CstZone := common.CstZone()
	fmt.Println("时区-",CstZone)

	//东八区时间
	DefaultTime := common.DefaultTime()
	fmt.Println("东八区时间-默认样式-",DefaultTime)

	//当前时间-有样式
	Now := common.Now()
	fmt.Println("当前时间-有样式-",Now)

	//当前日期-精确到天-有样式
	NowDate1 := common.NowDate(1)
	fmt.Println("当前日期-精确到天-有样式1-",NowDate1)

	//当前日期-精确到天-有样式
	NowDate2 := common.NowDate(2)
	fmt.Println("当前日期-精确到天-有样式2-",NowDate2)

	P.P("success")
}

/*
docker exec go_c go run demo/import_use.go
*/