package basedir

import (
	"fmt"
)

// CstAccount returns base directory path for customer accounts.
func CstAccount(category string) string {
	return fmt.Sprintf("cst_acc/%s", category)
}
