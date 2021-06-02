package helper

import (
	"strconv"
	"strings"
	"time"
	"unicode"
)

// UcFirst returns str with its first letter in uppercase.
func UcFirst(str string) string {
	runes := []rune(str)
	for i, v := range runes {
		return string(unicode.ToUpper(v)) + string(runes[i+1:])
	}
	return ""
}

// IsNumericString checks if a string contains only numbers.
func IsNumericString(str string) bool {
	for _, c := range str {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

// IsSQLDateString checks if a string is a valid SQL date.
func IsSQLDateString(s string) (isValidFormat, isValidDate bool) {
	parts := strings.Split(s, "-")
	if len(parts) == 3 {
		if len(parts[0]) == 4 && IsNumericString(parts[0]) &&
			len(parts[1]) == 2 && IsNumericString(parts[1]) &&
			len(parts[2]) == 2 && IsNumericString(parts[2]) {
			isValidFormat = true
			if _, err := time.Parse("2006-01-02", s); err == nil {
				isValidDate = true
			}
		}
	}
	return
}

// IntToString returns n as string.
func IntToString(n int) string {
	return strconv.FormatInt(int64(n), 10)
}

// Int32ToString returns n as string.
func Int32ToString(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}

// Int64ToString returns n as string.
func Int64ToString(n int64) string {
	return strconv.FormatInt(n, 10)
}

// StringToInt returns s as int.
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

// StringToInt64 returns s as int64.
func StringToInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

// StringToFloat32 returns s as float32.
func StringToFloat32(s string) (float32, error) {
	f, err := strconv.ParseFloat(s, 32)
	return float32(f), err
}

// StringToFloat64 returns s as float64.
func StringToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// SplitStringByLength returns an array of separated strings, each having the maximum length allowed.
func SplitStringByLength(str string, max int) []string {
	runes := []rune(str)
	runeLen := len(runes)
	if runeLen <= max {
		return []string{str}
	}
	count := runeLen / max
	if runeLen%max != 0 {
		count++
	}
	arr := make([]string, count)
	for i := 0; i < count; i++ {
		start := i * max
		end := start + max
		if end > runeLen {
			end = runeLen
		}
		arr[i] = string(runes[start:end])
	}
	return arr
}

// PadLeft adds padding string to the left of a string until the padded string reach the specified max length.
func PadLeft(str, pad string, maxLength int) string {
	if pad != "" {
		for len(str)+len(pad) <= maxLength {
			str = pad + str
		}
	}
	return str
}

// PadRight adds padding string to the right of a string until the padded string reach the specified max length.
func PadRight(str, pad string, maxLength int) string {
	if pad != "" {
		for len(str)+len(pad) <= maxLength {
			str = str + pad
		}
	}
	return str
}

// WithThousandSeparator formats a number with thousand separator. Make sure the number is not a negative value.
func WithThousandSeparator(number int64) string {
	if number < 0 {
		number = -number
	}
	s := Int64ToString(number)
	if number < 1000 {
		return s
	}
	ex := len(s) % 3
	if ex == 0 {
		ex = 3
	}
	res := s[:ex]
	for i := ex; i < len(s); i += 3 {
		res += "." + s[i:i+3]
	}
	return res
}

// GetFileExtension returns file's extension of a filename, that is the string after the last dot character.
func GetFileExtension(filename string) string {
	parts := strings.Split(filename, ".")
	if len(parts) > 1 {
		return "." + parts[len(parts)-1]
	}
	return ""
}
