package helper

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
)

// Defines phone number's min & max length.
const (
	PhoneMinLength = 7
	PhoneMaxLength = 15
)

// Defines password's min & max length.
const (
	PasswordMinLength = 6
	PasswordMaxLength = 32
)

// ValidatePhoneFormat checks if a phone number's format is valid.
func ValidatePhoneFormat(countryCallingCode, phone string) error {
	if phone == "" {
		return errors.New("Phone number is empty")
	} else if !IsNumericString(phone) {
		return errors.New("Phone number must be numeric")
	} else if strings.HasPrefix(phone, "0") {
		return errors.New("Phone number can't start with '0'")
	}
	lenCC, lenPhone := len(countryCallingCode), len(phone)
	min, max := PhoneMinLength-lenCC, PhoneMaxLength-lenCC
	if lenPhone < min || lenPhone > max {
		return fmt.Errorf("Phone number's length must be %d-%d digits", min, max)
	}
	return nil
}

// ValidateEmailFormat checks if an email's format is valid.
func ValidateEmailFormat(email string) error {
	if email == "" {
		return errors.New("Email address is empty")
	}
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if re.MatchString(email) {
		// TODO: Temporary. Check the domain.
		if strings.HasSuffix(strings.ToLower(email), "@mailinator.com") {
			return errors.New("Email address is invalid")
		}
		return nil
	}
	return errors.New("Email address format is invalid")
}

// ValidatePasswordFormat checks if a password's format is valid.
func ValidatePasswordFormat(password string) error {
	if password == "" {
		return errors.New("Password is empty")
	} else if l := len(password); l < PasswordMinLength || l > PasswordMaxLength {
		return fmt.Errorf("Password's length must be %d-%d characters", PasswordMinLength, PasswordMaxLength)
	}
	countUpper, countLower, countNumber, countSpecial := 0, 0, 0, 0
	for _, c := range password {
		if c >= '0' && c <= '9' {
			countNumber++
		} else if c >= 'a' && c <= 'z' {
			countLower++
		} else if c >= 'A' && c <= 'Z' {
			countUpper++
		} else {
			countSpecial++
		}
	}
	if countUpper == 0 || countLower == 0 || countNumber == 0 || countSpecial == 0 {
		return errors.New("Password must contain at least 1 lowercase, uppercase, and special characters and 1 number")
	}
	return nil
}

// ValidateIPAddressFormat checks if an IP address format is valid.
func ValidateIPAddressFormat(value string) (bool, error) {
	if value == "" {
		return false, errors.New("empty")
	}
	list := strings.Split(value, ",")
	for _, s := range list {
		s = strings.Trim(s, " ")
		if s == "" {
			return false, errors.New("item empty")
		}
		parts := strings.Split(s, ".")
		if len(parts) != 4 {
			return false, errors.New("parts count invalid")
		}
		for _, p := range parts {
			if p == "" || len(p) > 3 {
				return false, errors.New("part length invalid")
			} else if !IsNumericString(p) {
				return false, errors.New("part not number")
			}
			n, err := StringToInt(p)
			if err != nil {
				return false, errors.New("part not number")
			} else if n < 0 || n > 255 {
				return false, errors.New("part range invalid")
			}
		}
	}
	return true, nil
}

// CalculateAmountFromPercentageInt64 calculates amount from percentage.
func CalculateAmountFromPercentageInt64(baseAmount int64, percentage float64) int64 {
	if percentage == 0 {
		return 0
	}
	return int64(math.Round(float64(baseAmount) * percentage / 100))
}
