package model

import (
	"strconv"
	"strings"
	"time"
)

func StrToDate(date string) (time.Time, error) {
	split := strings.Split(date, "-")
	yearStr := split[0]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return time.Time{}, err
	}
	mouthStr := split[1]
	mouth, err := strconv.Atoi(mouthStr)
	if err != nil {
		return time.Time{}, err
	}
	dayStr := split[2]
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(year, time.Month(mouth), day, 0, 0, 0, 0, time.UTC), nil
}
