package dao

import "fmt"

// argPlaceholder helps with creating PostgreSQL argument placeholder.
type argPlaceholder struct {
	lastIndex int
}

// ResetIndex resets the last index to 0.
func (ap *argPlaceholder) ResetIndex() {
	ap.lastIndex = 0
}

// NextPlaceholder returns placeholder string with incremented index for each call.
func (ap *argPlaceholder) NextPlaceholder() string {
	ap.lastIndex = ap.lastIndex + 1
	return fmt.Sprintf("$%d", ap.lastIndex)
}
