package utils

import "time"

const (
	DatetimeFormart = "2006-01-02 15:04:05"
	DateFormart = "2006-01-02"
)

func GetNowTimezone() *time.Location {
	return time.FixedZone("CST", 3600 * 8)
}

func ConvertTimezone(now time.Time) time.Time{
	return now.In(GetNowTimezone())
}
