package model

import (
	"fmt"
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
	monthStr := split[1]
	mouth, err := strconv.Atoi(monthStr)
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

func StrDateSplit(date string) (int, int, int, error) {
	split := strings.Split(date, "-")
	yearStr := split[0]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return 0, 0, 0, err
	}
	monthStr := split[1]
	month, err := strconv.Atoi(monthStr)
	if err != nil {
		return 0, 0, 0, err
	}
	dayStr := split[2]
	day, err := strconv.Atoi(dayStr)
	if err != nil {
		return 0, 0, 0, err
	}
	return day, month, year, nil
}

func StrDateToMouthYear(date string) (int, int, error) {
	split := strings.Split(date, "-")
	yearStr := split[0]
	year, err := strconv.Atoi(yearStr)
	if err != nil {
		return 0, 0, err
	}
	monthStr := split[1]
	mouth, err := strconv.Atoi(monthStr)
	if err != nil {
		return 0, 0, err
	}
	return mouth, year, nil
}

func DateToDMY(date time.Time) (string, string, string) {
	year := fmt.Sprintf("%d", date.Year())
	month := fmt.Sprintf("%02d", date.Month())
	day := fmt.Sprintf("%d", date.Day())
	return day, month, year
}

func DateToDMYString(date time.Time) string {
	year, month, day := DateToDMY(date)
	return fmt.Sprintf("%s-%s-%s", year, month, day)
}
