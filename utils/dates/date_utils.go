package dates

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

// GetNow Util to retrieve current standard time
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString Util to retrieve current standard time string
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
