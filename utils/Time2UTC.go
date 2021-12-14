package utils

import "time"

func ConvertStringToTime(s string) (time.Time, error) {
	var loc, _ = time.LoadLocation("Asia/Ho_Chi_Minh")
	return time.ParseInLocation("2006-01-02 15:04:05", s, loc)
}
