package datetime

import "time"

func GetUnixTime() int64 {
	return time.Now().Unix()
}

// get current time,format: 2020-10-01 01:32:28
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
