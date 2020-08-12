package dates

import "time"

const (
	apiDateLayout   = "2006-01-02T15:04:05Z"
	apiDateDBLayout = "2006-01-02 15:04:05"
)

// GetNow Util to retrieve current standard time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString Util to retrieve current standard time string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowStringDB Util to retrieve current standard time string with DB Layout
func GetNowStringDB() string {
	return GetNow().Format(apiDateDBLayout)
}
