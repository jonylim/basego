package helper

import (
	"fmt"
	"strings"
	"testing"
)

func TestUcFirst(t *testing.T) {
	var tests = []struct {
		input, expected string
	}{
		{"", ""},
		{"hello world", "Hello world"},
		{"!important", "!important"},
		{"éCafé", "ÉCafé"},
	}
	for _, test := range tests {
		if out := UcFirst(test.input); out != test.expected {
			t.Errorf(`UcFirst("%v") = "%v"; expected "%v"`, test.input, out, test.expected)
		}
	}
}

func TestIsNumericString(t *testing.T) {
	var tests = []struct {
		input    string
		expected bool
	}{
		{"081234567890", true},
		{" 081234567890", false},
		{"0812-3456-7890", false},
		{"0812.3456.7890", false},
		{"0812 3456 7890", false},
		{"+6281234567890", false},
	}
	for _, test := range tests {
		if out := IsNumericString(test.input); out != test.expected {
			t.Errorf(`IsNumericString("%v") = %v; expected %v`, test.input, out, test.expected)
		}
	}
}

func TestIsSQLDateString(t *testing.T) {
	var tests = []struct {
		input               string
		expectedValidFormat bool
		expectedValidDate   bool
	}{
		{"2006-01-02", true, true},
		{"2006-01-02 15:04:05", false, false},
		{"2006-01-2", false, false},
		{"2006-1-02", false, false},
		{"06-01-02", false, false},
		{"2020-01-32", true, false},
		{"2020-02-29", true, true},
		{"2000-02-29", true, true},
		{"2100-02-29", true, false},
		{"yyyy-MM-dd", false, false},
	}
	for _, test := range tests {
		if out1, out2 := IsSQLDateString(test.input); out1 != test.expectedValidFormat || out2 != test.expectedValidDate {
			t.Errorf(`IsSQLDateString("%s") = (%v, %v); expected (%v, %v)`, test.input, out1, out2, test.expectedValidFormat, test.expectedValidDate)
		}
	}
}

func TestIntToString(t *testing.T) {
	var tests = []struct {
		input    int
		expected string
	}{
		{-2147483647, "-2147483647"},
		{2147483647, "2147483647"},
		{4294967295, "4294967295"},
		{-4294967295, "-4294967295"},
		{81234567890, "81234567890"},
		{81234567890, "81234567890"},
		{9223372036854775807, "9223372036854775807"},
		{-9223372036854775807, "-9223372036854775807"},
	}
	for _, test := range tests {
		if out := IntToString(test.input); out != test.expected {
			t.Errorf(`IntToString(%v) = "%v"; expected "%v"`, test.input, out, test.expected)
		}
	}
}

func TestInt32ToString(t *testing.T) {
	var tests = []struct {
		input    int32
		expected string
	}{
		{-2147483647, "-2147483647"},
		{2147483647, "2147483647"},
	}
	for _, test := range tests {
		if out := Int32ToString(test.input); out != test.expected {
			t.Errorf(`Int32ToString(%v) = "%v"; expected "%v"`, test.input, out, test.expected)
		}
	}
}

func TestInt64ToString(t *testing.T) {
	var tests = []struct {
		input    int64
		expected string
	}{
		{-2147483647, "-2147483647"},
		{2147483647, "2147483647"},
		{4294967295, "4294967295"},
		{-4294967295, "-4294967295"},
		{81234567890, "81234567890"},
		{81234567890, "81234567890"},
		{9223372036854775807, "9223372036854775807"},
		{-9223372036854775807, "-9223372036854775807"},
	}
	for _, test := range tests {
		if out := Int64ToString(test.input); out != test.expected {
			t.Errorf(`Int64ToString(%v) = "%v"; expected "%v"`, test.input, out, test.expected)
		}
	}
}

func TestStringToInt(t *testing.T) {
	var tests = []struct {
		input    string
		expected int
		sErr     string
	}{
		{"081234567890", 81234567890, "nil"},
		{"-2147483647", -2147483647, "nil"},
		{"2147483647", 2147483647, "nil"},
		{"4294967295", 4294967295, "nil"},
		{"-4294967295", -4294967295, "nil"},
		{"81234567890", 81234567890, "nil"},
		{"81234567890", 81234567890, "nil"},
		{"9223372036854775807", 9223372036854775807, "nil"},
		{"-9223372036854775807", -9223372036854775807, "nil"},
		{"24d", 0, "err"},
		{"24f", 0, "err"},
	}
	for _, test := range tests {
		out, err := StringToInt(test.input)
		serr := "nil"
		if err != nil {
			serr = "err"
		}
		if out != test.expected || serr != test.sErr {
			t.Errorf(`StringToInt("%v") = (%v, %v); expected (%v, %v); err: %v`, test.input, out, serr, test.expected, test.sErr, err)
		}
	}
}

func TestStringToInt64(t *testing.T) {
	var tests = []struct {
		input    string
		expected int64
		sErr     string
	}{
		{"081234567890", 81234567890, "nil"},
		{"-2147483647", -2147483647, "nil"},
		{"2147483647", 2147483647, "nil"},
		{"4294967295", 4294967295, "nil"},
		{"-4294967295", -4294967295, "nil"},
		{"81234567890", 81234567890, "nil"},
		{"81234567890", 81234567890, "nil"},
		{"9223372036854775807", 9223372036854775807, "nil"},
		{"-9223372036854775807", -9223372036854775807, "nil"},
		{"24d", 0, "err"},
		{"24f", 0, "err"},
	}
	for _, test := range tests {
		out, err := StringToInt64(test.input)
		serr := "nil"
		if err != nil {
			serr = "err"
		}
		if out != test.expected || serr != test.sErr {
			t.Errorf(`StringToInt64("%v") = (%v, %v); expected (%v, %v); err: %v`, test.input, out, serr, test.expected, test.sErr, err)
		}
	}
}

func TestStringToFloat32(t *testing.T) {
	var tests = []struct {
		input    string
		expected float32
		sErr     string
	}{
		{"123.0", 123.0, "nil"},
		{"081234.567890", 81234.567890, "nil"},
		{"-21.47483647", -21.47483647, "nil"},
		{"2.147483647", 2.147483647, "nil"},
		{"4.294967295", 4.294967295, "nil"},
		{"-42.94967295", -42.94967295, "nil"},
		{"812.34567890", 812.34567890, "nil"},
		{"8123456789.0", 8123456789.0, "nil"},
		{"92233720368547758.07", 92233720368547758.07, "nil"},
		{"-9.223372036854775807", -9.223372036854775807, "nil"},
		{"24d", 0, "err"},
		{"24f", 0, "err"},
	}
	for _, test := range tests {
		out, err := StringToFloat32(test.input)
		serr := "nil"
		if err != nil {
			serr = "err"
		}
		if out != test.expected || serr != test.sErr {
			t.Errorf(`StringToFloat32("%v") = (%v, %v); expected (%v, %v); err: %v`, test.input, out, serr, test.expected, test.sErr, err)
		}
	}
}

func TestStringToFloat64(t *testing.T) {
	var tests = []struct {
		input    string
		expected float64
		sErr     string
	}{
		{"123.0", 123.0, "nil"},
		{"081234.567890", 81234.567890, "nil"},
		{"-21.47483647", -21.47483647, "nil"},
		{"2.147483647", 2.147483647, "nil"},
		{"4.294967295", 4.294967295, "nil"},
		{"-42.94967295", -42.94967295, "nil"},
		{"812.34567890", 812.34567890, "nil"},
		{"8123456789.0", 8123456789.0, "nil"},
		{"92233720368547758.07", 92233720368547758.07, "nil"},
		{"-9.223372036854775807", -9.223372036854775807, "nil"},
		{"24d", 0, "err"},
		{"24f", 0, "err"},
	}
	for _, test := range tests {
		out, err := StringToFloat64(test.input)
		serr := "nil"
		if err != nil {
			serr = "err"
		}
		if out != test.expected || serr != test.sErr {
			t.Errorf(`StringToFloat64("%v") = (%v, %v); expected (%v, %v); err: %v`, test.input, out, serr, test.expected, test.sErr, err)
		}
	}
}

func TestSplitStringByLength(t *testing.T) {
	var tests = []struct {
		inputStr string
		inputMax int
		expected []string
	}{
		{
			inputStr: "Lorem ipsum dolor sit amet",
			inputMax: 30,
			expected: []string{
				"Lorem ipsum dolor sit amet",
			},
		},
		{
			inputStr: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			inputMax: 20,
			expected: []string{
				"Lorem ipsum dolor si",
				"t amet, consectetur ",
				"adipiscing elit, sed",
				" do eiusmod tempor i",
				"ncididunt ut labore ",
				"et dolore magna aliq",
				"ua.",
			},
		},
		{
			inputStr: "Lorém ipsum dolor sit amét, conséctétur adipiscing élit, séd do éiusmod témpor incididunt ut laboré ét doloré magna aliq",
			inputMax: 12,
			expected: []string{
				"Lorém ipsum ",
				"dolor sit am",
				"ét, consécté",
				"tur adipisci",
				"ng élit, séd",
				" do éiusmod ",
				"témpor incid",
				"idunt ut lab",
				"oré ét dolor",
				"é magna aliq",
			},
		},
	}
	for _, test := range tests {
		out := SplitStringByLength(test.inputStr, test.inputMax)
		lenOut, lenExpected := len(out), len(test.expected)
		if lenOut != lenExpected {
			t.Errorf(`len(SplitStringByLength("%v", %v)) = %v; expected %v`, test.inputStr, test.inputMax, lenOut, lenExpected)
		} else {
			failed := false
			var sb strings.Builder
			sb.WriteString("[")
			for i := range out {
				if out[i] != test.expected[i] {
					failed = true
					sb.WriteString(fmt.Sprintf("\n    \"%s\"; expected \"%s\"", out[i], test.expected[i]))
				} else {
					sb.WriteString(fmt.Sprintf("\n    \"%s\"", out[i]))
				}
			}
			sb.WriteString("\n]")
			if failed {
				t.Errorf(`SplitStringByLength("%v", %v) = %s`, test.inputStr, test.inputMax, sb.String())
			}
		}
	}
}

func TestGetFileExtension(t *testing.T) {
	var tests = []struct {
		input, expected string
	}{
		{"", ""},
		{"some-file", ""},
		{"hello.jpg", ".jpg"},
		{"my.name.is.png", ".png"},
		{"éCafé.mp4", ".mp4"},
	}
	for _, test := range tests {
		if out := GetFileExtension(test.input); out != test.expected {
			t.Errorf(`GetFileExtension("%v") = "%v"; expected "%v"`, test.input, out, test.expected)
		}
	}
}
