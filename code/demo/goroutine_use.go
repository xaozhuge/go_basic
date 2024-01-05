package main

import (
	"fmt"
	"sync"
	"code/P"
	"time"
)

var wg sync.WaitGroup


func download(url string) {
	fmt.Println("start to download", url)
	P.D(time.Second)
	time.Sleep(time.Second) // 模拟耗时操作
	wg.Done()
}

func main() {
	startTime := time.Now()
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go download("test.com/" + string(i+'0'))
	}
	wg.Wait()
	fmt.Println("Done!")
	elapsedTime := time.Since(startTime)
	P.D(elapsedTime)

}



/*
docker exec go_c go run demo/goroutine_use.go
*/

/*
1. sync.WaitGroup 是一个用于等待一组 goroutine 完成执行的同步工具。
它通常用于确保所有的 goroutine 都执行完毕后再继续执行下一步。
2. sync.WaitGroup 通常包含了一个内部计数器，可以通过调用 Add 方法增加计数，
通过调用 Done 方法减少计数，以及通过调用 Wait 方法等待计数器变为零。
这样，可以在主 goroutine 中等待所有的子 goroutine 完成执行，
然后再继续执行主 goroutine 的后续逻辑。
3. go download()：启动新的协程并发执行 download 函数。
4. wg.Wait()：等待所有的协程执行结束。
5. 可以看到串行需要 3s 的下载操作，并发后，只需要 1s。





 
 */