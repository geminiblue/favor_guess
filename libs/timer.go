package libs

import (
	"time"
)

const FORMAT = "2006-01-02 15:04:05"

func getTimeByDuration(duration int, format string) string {
	currentTime := time.Now()
	ti := currentTime.AddDate(0, 0, duration)
	return ti.Format(format)
}

func Now() string {
	return time.Now().Format(FORMAT)
}

func NowTime() time.Time {
	return time.Now()
}

//获取之前或者之后的时间
func GetTimeByDurationBegin(duration int) string {
	format := "2006-01-02 00:00:00"
	return getTimeByDuration(duration, format)
}

//获取之前或者之后的时间
func GetTimeByDurationEnd(duration int) string {
	format := "2006-01-02 23:59:59"
	return getTimeByDuration(duration, format)
}

//字符串时间转时间戳
func StrToTime(str string) time.Time {
	p, _ := time.Parse(FORMAT, str)
	return p
}

//获取当前时间戳
func TimeStamp() int64 {
	return time.Now().Unix()
}
