package utils

import "time"

func GetCurrentTimestampMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
