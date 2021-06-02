package helper

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// UnixMillisecond returns t as a Unix time, the number of milliseconds elapsed since January 1, 1970 UTC.
func UnixMillisecond(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// UnixMicrosecond returns t as a Unix time, the number of microseconds elapsed since January 1, 1970 UTC.
func UnixMicrosecond(t time.Time) int64 {
	return t.UnixNano() / 1e3
}

// FromUnixMillisecond returns Unix milliseconds as time.Time.
func FromUnixMillisecond(millis int64) time.Time {
	return time.Unix(0, millis*1e6)
}

// GetLastDayInMonth returns last day in a specified year and month.
func GetLastDayInMonth(year, month int) (lastDay int, err error) {
	var t time.Time
	t, err = time.Parse("2006-1-2", fmt.Sprintf("%d-%d-1", year, month))
	if err == nil {
		t = t.AddDate(0, 1, 0).AddDate(0, 0, -1)
		lastDay = t.Day()
	}
	return
}

// AddMonth adds n month to t. If the resulting month has less days than t, returns the last day in the added month.
func AddMonth(t time.Time, n int) time.Time {
	anchorDay := t.Day()
	firstOfMonth := t
	if t.Day() != 1 {
		firstOfMonth = t.AddDate(0, 0, 1-t.Day())
	}
	res := firstOfMonth.AddDate(0, n, 0)
	if lastDay, _ := GetLastDayInMonth(res.Year(), int(res.Month())); anchorDay > lastDay {
		res = res.AddDate(0, 0, lastDay-1)
	} else {
		res = res.AddDate(0, 0, anchorDay-1)
	}
	return res
}

// AddYear adds n year to t. If the resulting month has less days than t, returns the last day in the added month & year.
func AddYear(t time.Time, n int) time.Time {
	anchorDay := t.Day()
	firstOfMonth := t
	if t.Day() != 1 {
		firstOfMonth = t.AddDate(0, 0, 1-t.Day())
	}
	res := firstOfMonth.AddDate(n, 0, 0)
	if lastDay, _ := GetLastDayInMonth(res.Year(), int(res.Month())); anchorDay > lastDay {
		res = res.AddDate(0, 0, lastDay-1)
	} else {
		res = res.AddDate(0, 0, anchorDay-1)
	}
	return res
}

// ParseTimeHHMM parses time in "HH:mm" format to time.Duration.
func ParseTimeHHMM(t string) (time.Duration, error) {
	parts := strings.Split(t, ":")
	if len(parts) != 2 {
		return 0, errors.New("Invalid time format")
	}
	hours, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}
	mins, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}
	if hours < 0 || hours > 23 || mins < 0 || mins > 59 {
		return 0, fmt.Errorf("Invalid time: %s", t)
	}
	return time.Duration(hours)*time.Hour + time.Duration(mins)*time.Minute, nil
}

// ValidateTimeFormatHHMM validates time in "HH:mm" format.
func ValidateTimeFormatHHMM(time string) error {
	parts := strings.Split(time, ":")
	if len(time) != 5 || len(parts) != 2 {
		return errors.New("Invalid time format")
	} else if len(parts[0]) != 2 || len(parts[1]) != 2 {
		return errors.New("Invalid time format")
	} else if !IsNumericString(parts[0]) || !IsNumericString(parts[1]) {
		return errors.New("Invalid time format")
	}
	return nil
}

var mMonthRomans = map[int]string{
	1:  "I",
	2:  "II",
	3:  "III",
	4:  "IV",
	5:  "V",
	6:  "VI",
	7:  "VII",
	8:  "VIII",
	9:  "IX",
	10: "X",
	11: "XI",
	12: "XII",
}

// GetRomanNumeralByMonth returns Roman numeral for a month.
func GetRomanNumeralByMonth(month int) string {
	return mMonthRomans[month]
}

var mMonthShortNames = map[int]string{
	1:  "Jan",
	2:  "Feb",
	3:  "Mar",
	4:  "Apr",
	5:  "May",
	6:  "Jun",
	7:  "Jul",
	8:  "Aug",
	9:  "Sep",
	10: "Oct",
	11: "Nov",
	12: "Dec",
}

// GetMonthShortName returns short name for a month.
func GetMonthShortName(month int) string {
	return mMonthShortNames[month]
}
