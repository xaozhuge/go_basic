package P

import "os"
import "fmt"

func P(temp_var ...interface{}) {
	// 使用fmt.Println输出
	fmt.Println(temp_var...)
	// 然后退出程序
	os.Exit(0)
}