package common

import "time"

// 时区选择
func CstZone() *time.Location {
    return time.FixedZone("CST", 8*3600)
}

//东八区的时间
func DefaultTime() time.Time {
	cst_zone := CstZone()
	return time.Now().In(cst_zone)
}

// 当前时间
func Now() string {
    now_time := DateFormat(DefaultTime(), 3) 
    return now_time
}

// 日期的不同格式
func DateFormat(t time.Time, formatType int) string {
	var format string

	switch formatType {
		case 1:
			format = "2006-01-02"
		case 2:
			format = "20060102"
		case 3:
			format = "2006-01-02 15:04:05"
		default:
			format = "2006-01-02"
	}

	return t.Format(format)
}

// 天的不同格式
func NowDate(formatType int) string {
    now_date := DateFormat(DefaultTime(), formatType) 
    return now_date
}

