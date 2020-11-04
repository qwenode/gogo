package datetime

import "time"

// GetUnixTime get current unix time
func GetUnixTime() int64 {
	return time.Now().Unix()
}

// GetCurrentTime get current time,format: 2020-10-01 01:32:28
func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetBeginOfTheDayByInt64 convert 2020-10-10 21:22:33 to 2020-10-10 00:00:00
func GetBeginOfTheDayByInt64(unix int64) int64 {
	if unix < 100000000 {
		return 0
	}
	format := time.Unix(unix, 0).Format("2006-01-02")
	parse, _ := time.ParseInLocation("2006-01-02", format, time.Local)
	return parse.Unix()
}

// GetBeginOfTheDayByInt convert 2020-10-10 21:22:33 to 2020-10-10 00:00:00
func GetBeginOfTheDayByInt(unix int) int {
	return int(GetBeginOfTheDayByInt64(int64(unix)))
}

// GetEndOfTheDayByInt64 convert 2020-10-10 21:22:33 to 2020-10-10 23:59:59
func GetEndOfTheDayByInt64(unix int64) int64 {
	if unix < 100000000 {
		return 0
	}
	return GetBeginOfTheDayByInt64(unix) + 86399
}

// GetEndOfTheDayByInt convert 2020-10-10 21:22:33 to 2020-10-10 23:59:59
func GetEndOfTheDayByInt(unix int) int {
	return int(GetEndOfTheDayByInt64(int64(unix)))
}
