package helper

// StringExists checks if a string exists in an array of strings.
func StringExists(items []string, find string) bool {
	for _, it := range items {
		if it == find {
			return true
		}
	}
	return false
}

// IntExists checks if an int exists in an array of ints.
func IntExists(items []int, find int) bool {
	for _, it := range items {
		if it == find {
			return true
		}
	}
	return false
}

// Int32Exists checks if an int32 exists in an array of int32s.
func Int32Exists(items []int32, find int32) bool {
	for _, it := range items {
		if it == find {
			return true
		}
	}
	return false
}

// Int64Exists checks if an int64 exists in an array of int64s.
func Int64Exists(items []int64, find int64) bool {
	for _, it := range items {
		if it == find {
			return true
		}
	}
	return false
}
