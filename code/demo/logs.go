package main

import (
    "fmt"
    "os"
    "io"
    "time"
)

func logs(log_content string, filename string) {
    var file *os.File
    var err error
    if checkFileIsExist(filename) { //如果文件存在
        file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
        fmt.Println("文件存在")
    } else {//文件不存在，创建文件
        file, err = os.Create(filename) 
        fmt.Println("文件不存在")
    }
    checkErr(err)
    content_length, err := io.WriteString(file, log_content+"\n") //写入文件(字符串)
    checkErr(err)
    fmt.Printf("写入 %d 个字节\n", content_length)
}

/**
 * 判断err
 */
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) bool {
    var exist = true
    if _, err := os.Stat(filename); os.IsNotExist(err) {
        exist = false
    }
    return exist
}

func main() {
    var cst_zone = time.FixedZone("CST", 8 * 3600) // 东八
    startTime := time.Now().In(cst_zone).Format("2006-01-02 15:04:05")
    var log_content = startTime+":\n"+"日志内容"
    var filename    = "test.log"
    logs(log_content, filename)
}
