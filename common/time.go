package common

import (
	"errors"
	"github.com/oneliang/util-golang/constants"
	"time"
)

func GetZeroTime(milliSecond int64, modulusTime int64) int64 {
	var currentTime = milliSecond //time zone is zero +0000
	var timeZoneMilliSecondOffset = GetTimeZoneMilliSecondOffset()
	var currentTimeZoneTime = currentTime + int64(timeZoneMilliSecondOffset)
	var retainTime = currentTimeZoneTime % modulusTime //current time zone time
	return currentTime - retainTime                    //recovery to 0 time zone
}

func GetTimeZoneMilliSecondOffset() int {
	_, secondOffset := time.Now().Zone()
	return secondOffset * constants.TIME_MILLISECONDS_OF_SECOND
}
func GetTimeZoneInt() int {
	return GetTimeZoneMilliSecondOffset() / constants.TIME_MILLISECONDS_OF_HOUR
}

func GetDayZeroTime(millisecond int64) int64 {
	return GetZeroTime(millisecond, constants.TIME_MILLISECONDS_OF_DAY)
}

func GetHourZeroTime(millisecond int64) int64 {
	return GetZeroTime(millisecond, constants.TIME_MILLISECONDS_OF_HOUR)
}

func GetMinuteZeroTime(millisecond int64) int64 {
	return GetZeroTime(millisecond, constants.TIME_MILLISECONDS_OF_MINUTE)
}

func GetSecondZeroTime(millisecond int64) int64 {
	return GetZeroTime(millisecond, constants.TIME_MILLISECONDS_OF_SECOND)
}

func GetZeroTimeNext(millisecond int64, modulusTime int64, offset int) int64 {
	return GetZeroTime(millisecond, modulusTime) + int64(offset)*modulusTime
}

func GetDayZeroTimeNext(millisecond int64, offset int) int64 {
	return GetDayZeroTime(millisecond) + int64(offset)*constants.TIME_MILLISECONDS_OF_DAY
}

func GetHourZeroTimeNext(millisecond int64, offset int) int64 {
	return GetHourZeroTime(millisecond) + int64(offset)*constants.TIME_MILLISECONDS_OF_HOUR
}

func GetMinuteZeroTimeNext(millisecond int64, offset int) int64 {
	return GetMinuteZeroTime(millisecond) + int64(offset)*constants.TIME_MILLISECONDS_OF_MINUTE
}

func GetSecondZeroTimeNext(millisecond int64, offset int) int64 {
	return GetSecondZeroTime(millisecond) + int64(offset)*constants.TIME_MILLISECONDS_OF_SECOND
}

func GetZeroTimePrevious(millisecond int64, modulusTime int64, offset int) int64 {
	return GetZeroTime(millisecond, modulusTime) - int64(offset)*modulusTime
}

func GetDayZeroTimePrevious(millisecond int64, offset int) int64 {
	return GetDayZeroTime(millisecond) - int64(offset)*constants.TIME_MILLISECONDS_OF_DAY
}

func GetHourZeroTimePrevious(millisecond int64, offset int) int64 {
	return GetHourZeroTime(millisecond) - int64(offset)*constants.TIME_MILLISECONDS_OF_HOUR
}

func GetMinuteZeroTimePrevious(millisecond int64, offset int) int64 {
	return GetMinuteZeroTime(millisecond) - int64(offset)*constants.TIME_MILLISECONDS_OF_MINUTE
}

func GetSecondZeroTimePrevious(millisecond int64, offset int) int64 {
	return GetSecondZeroTime(millisecond) - int64(offset)*constants.TIME_MILLISECONDS_OF_MINUTE
}

func GetCurrentMonthDatesOffset(monthOffset int, monthCount int) (startTime, endTime time.Time, err error) {
	if monthCount <= 0 {
		return time.Time{}, time.Time{}, errors.New("parameter month count must be bigger than zero")
	}
	now := time.Now()
	startOfMonth := now.AddDate(0, monthOffset, -now.Day()+1)
	startDate := time.Date(startOfMonth.Year(), startOfMonth.Month(), startOfMonth.Day(), 0, 0, 0, 0, time.Local)
	endOfMonth := now.AddDate(0, monthOffset+monthCount, -now.Day())
	endDate := time.Date(endOfMonth.Year(), endOfMonth.Month(), endOfMonth.Day(), 23, 59, 59, 999999999, time.Local)

	return startDate, endDate, nil
}
