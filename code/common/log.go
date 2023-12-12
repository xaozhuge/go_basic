package common

import (
    "fmt"
    "encoding/json"

)

func Logs(filename string, log_var ...string) {
    NowDate := NowDate(2)
    Now := Now()
    var path string
    if len(filename) == 0 {
        path = "log/" + NowDate+  ".log"
    }else{
        path = "log/" + filename + "-" + NowDate + ".log"
    }

    log_data := make(map[string]string)
    log_data["time"] = Now

    for i, v := range log_var {
        key := fmt.Sprintf("var%d", i+1)
        log_data[key] = v
    }

    jsonData, _ := json.Marshal(log_data)
    content := string(jsonData)
    WriteFile(path, content)
}


