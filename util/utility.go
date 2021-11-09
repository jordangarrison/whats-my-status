package utility

import (
	"time"

	"github.com/hako/durafmt"
)

func GetEpochTime(hrTime string) (int64, error) {
	duration, err := durafmt.ParseString(hrTime)
	if err != nil {
		return -1, err
	}
	seconds := int64(duration.Duration().Seconds())
	return time.Now().Unix() + seconds, nil
}

func GetISO8601Time(hrTime string) (string, error) {
	duration, err := durafmt.ParseString(hrTime)
	if err != nil {
		return "", err
	}
	seconds := int64(duration.Duration().Seconds())
	return time.Now().Add(time.Duration(seconds) * time.Second).Format(time.RFC3339), nil
}
