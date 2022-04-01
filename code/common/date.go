package common
import "time"
func Now() string {
	// 东八区
    var cst_zone = time.FixedZone("CST", 8 * 3600)
    now_time := time.Now().In(cst_zone).Format("2006-01-02 15:04:05")
    return now_time
}
