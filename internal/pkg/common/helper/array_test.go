package helper

import (
	"strings"
	"testing"
)

func TestStringExists(t *testing.T) {
	items := []string{"Apple", "Orange", "Avocade", "Mango", "Banana"}
	var tests = []struct {
		input    string
		expected bool
	}{
		{"", false},
		{"apple", false},
		{"Orang", false},
		{"Orange", true},
	}
	printItems := true
	for _, test := range tests {
		if out := StringExists(items, test.input); out != test.expected {
			if printItems {
				t.Logf(`items := [ "%v" ]`, strings.Join(items, `", "`))
				printItems = false
			}
			t.Errorf(`StringExists(items, "%v") = %v; expected %v`, test.input, out, test.expected)
		}
	}
}

func TestIntExists(t *testing.T) {
	items := []int{20, -99, 123, 0, 777}
	var tests = []struct {
		input    int
		expected bool
	}{
		{-1, false},
		{1, false},
		{-20, false},
		{20, true},
		{-99, true},
		{99, false},
		{-0, true},
		{0, true},
		{123, true},
		{777, true},
	}
	printItems := true
	for _, test := range tests {
		if out := IntExists(items, test.input); out != test.expected {
			if printItems {
				t.Logf(`items := %v`, items)
				printItems = false
			}
			t.Errorf(`IntExists(items, %v) = %v; expected %v`, test.input, out, test.expected)
		}
	}
}

func TestInt32Exists(t *testing.T) {
	items := []int32{20, -99, 123, 0, 777}
	var tests = []struct {
		input    int32
		expected bool
	}{
		{-1, false},
		{1, false},
		{-20, false},
		{20, true},
		{-99, true},
		{99, false},
		{-0, true},
		{0, true},
		{123, true},
		{777, true},
	}
	printItems := true
	for _, test := range tests {
		if out := Int32Exists(items, test.input); out != test.expected {
			if printItems {
				t.Logf(`items := %v`, items)
				printItems = false
			}
			t.Errorf(`Int32Exists(items, %v) = %v; expected %v`, test.input, out, test.expected)
		}
	}
}

func TestInt64Exists(t *testing.T) {
	items := []int64{20, -99, 123, 0, 777}
	var tests = []struct {
		input    int64
		expected bool
	}{
		{-1, false},
		{1, false},
		{-20, false},
		{20, true},
		{-99, true},
		{99, false},
		{-0, true},
		{0, true},
		{123, true},
		{777, true},
	}
	printItems := true
	for _, test := range tests {
		if out := Int64Exists(items, test.input); out != test.expected {
			if printItems {
				t.Logf(`items := %v`, items)
				printItems = false
			}
			t.Errorf(`Int64Exists(items, %v) = %v; expected %v`, test.input, out, test.expected)
		}
	}
}
