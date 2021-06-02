package dao

import (
	"database/sql"
	"fmt"
)

// dao defines base struct for DAOs.
type dao struct {
	db          *sql.DB
	withDeleted bool
}

func (instance *dao) WithDeleted() {
	instance.withDeleted = true
}

// SQLRowOrRows represents sql.Row or sql.Rows.
type SQLRowOrRows interface {
	Scan(dest ...interface{}) error
}

func sqlTimestampToUnixMilliseconds(column string) string {
	return "COALESCE(FLOOR(EXTRACT(EPOCH FROM " + column + ") * 1000), 0)::BIGINT"
}

func sqlTimestampToUnixNanoseconds(column string) string {
	return "COALESCE(FLOOR(EXTRACT(EPOCH FROM " + column + ") * 1e9), 0)::BIGINT"
}

func sqlUnixMillisecondsToTimestamp(paramPlaceholder string) string {
	return fmt.Sprintf("TO_TIMESTAMP(%s::DECIMAL / 1000)", paramPlaceholder)
}

// Ascending or descending order.
const (
	AscendingOrder  = "ASC"
	DescendingOrder = "DESC"
)

func isValidOrder(v string) bool {
	switch v {
	case AscendingOrder, DescendingOrder:
		return true
	}
	return false
}
