package common

import (
    "os"
    "io"
)

func Logs(log_content string, filename string) int {
    var file *os.File
    var err error
    if checkFileIsExist(filename) { //如果文件存在
        file, err = os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
    } else {//文件不存在，创建文件
        file, err = os.Create(filename) 
    }
    checkErr(err)
    content_length, err := io.WriteString(file, log_content+"\n") //写入文件(字符串)
    checkErr(err)
    return content_length
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

