package common

import (
    "bufio"
    "fmt"
    "os"
)

func WriteFile(filePath string, content string) {
    // 打开文件以进行追加写入，如果文件不存在则创建
    file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // 创建带缓冲的写入器
    writer := bufio.NewWriter(file)

    // 写入内容
    _, err = writer.WriteString(content + "\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }

    // 刷新缓冲并确保数据被写入文件
    err = writer.Flush()
    if err != nil {
        fmt.Println("Error flushing writer:", err)
        return
    }
}
