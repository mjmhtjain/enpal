package util

import "time"

func UniversalTimeFormat(t time.Time) string {
	return time.Time(t).UTC().Format("2006-01-02T15:04:05.000Z")
}
