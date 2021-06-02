package helper

import (
	"strings"
	"testing"
)

func TestValidatePhoneFormat(t *testing.T) {
	var tests = []struct {
		inputCode, inputPhone string
		errKeyword            string
	}{
		{"62", "", "empty"},
		{"62", "81234567890", ""},
		{"", "12345", "length"},
		{"62", "12345678901234", "length"},
		{"", "1234567890123456", "length"},
		{"62", "812-3456-7890", "numeric"},
		{"62", "812 3456 7890", "numeric"},
		{"62", "081234567890", "'0'"},
	}
	for _, test := range tests {
		err := ValidatePhoneFormat(test.inputCode, test.inputPhone)
		if test.errKeyword == "" {
			if err != nil {
				t.Errorf(`ValidatePhoneFormat("%v", "%v") = error; expected nil`, test.inputCode, test.inputPhone)
			}
		} else {
			if err == nil {
				t.Errorf(`ValidatePhoneFormat("%v", "%v") = nil; expected error keyword "%s"`, test.inputCode, test.inputPhone, test.errKeyword)
			} else if !strings.Contains(err.Error(), test.errKeyword) {
				t.Errorf(`ValidatePhoneFormat("%v", "%v") = "%v"; expected error keyword "%s"`, test.inputCode, test.inputPhone, err, test.errKeyword)
			}
		}
	}
}

func TestValidateEmailFormat(t *testing.T) {
	var tests = []struct {
		input      string
		errKeyword string
	}{
		{"", "empty"},
		{"admin@localhost", ""},
		{"cs@example.com", ""},
		{"example.com", "invalid"},
		{"hello@world@example123.com", "invalid"},
		{"hello#world@example123.com", ""},
		{"hello.world@sub.example123.com", ""},
		{"hello@sub-example123.com", ""},
		{"hello.world@sub_example123.com", "invalid"},
		{"hello.world@.example.com", "invalid"},
		{"hello.world@123.example.com", ""},
		{"hello.world@example.com-", "invalid"},
		{"hello.world@-example.com", "invalid"},
	}
	for _, test := range tests {
		err := ValidateEmailFormat(test.input)
		if test.errKeyword == "" {
			if err != nil {
				t.Errorf(`ValidateEmailFormat("%v") = error; expected nil`, test.input)
			}
		} else {
			if err == nil {
				t.Errorf(`ValidateEmailFormat("%v") = nil; expected error keyword "%s"`, test.input, test.errKeyword)
			} else if !strings.Contains(err.Error(), test.errKeyword) {
				t.Errorf(`ValidateEmailFormat("%v") = "%v"; expected error keyword "%s"`, test.input, err, test.errKeyword)
			}
		}
	}
}

func TestValidatePasswordFormat(t *testing.T) {
	var tests = []struct {
		input      string
		errKeyword string
	}{
		{"", "empty"},
		{"12345", "length"},
		{"123456789012345678901234567890123", "length"},
		{"Aa@123", ""},
		{"Aa@123Aa@123Aa@123Aa@123Aa@123!?", ""},
		{"AA@123", "lowercase"},
		{"ab@123", "uppercase"},
		{"Abc123", "special"},
		{"A@bbcc", "number"},
	}
	for _, test := range tests {
		err := ValidatePasswordFormat(test.input)
		if test.errKeyword == "" {
			if err != nil {
				t.Errorf(`ValidatePasswordFormat("%v") = error; expected nil`, test.input)
			}
		} else {
			if err == nil {
				t.Errorf(`ValidatePasswordFormat("%v") = nil; expected error keyword "%s"`, test.input, test.errKeyword)
			} else if !strings.Contains(err.Error(), test.errKeyword) {
				t.Errorf(`ValidatePasswordFormat("%v") = "%v"; expected error keyword "%s"`, test.input, err, test.errKeyword)
			}
		}
	}
}

func TestValidateIPAddressFormat(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"123.123.123.123", true},
		{"0.1.2.3, 4.5.6.7, 8.9.10.11", true},
		{"123.123.123.123,", false},
		{"255.255.255.256", false},
		{"1.2.3, 5.6.7.8", false},
		{"1.2.3.4, 5.6.7.8.9", false},
		{"-1.1.1.1", false},
	}
	for _, test := range tests {
		if out, err := ValidateIPAddressFormat(test.input); out != test.expected {
			t.Errorf(`ValidateIPAddressFormat("%v") = %v; expected %v; err: %v`, test.input, out, test.expected, err)
		}
	}
}
