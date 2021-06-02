package repository

import "errors"

var (
	errNotFound = errors.New("repository: data not found")
	errDatabase = errors.New("repository: database error")
	errInternal = errors.New("repository: internal server error")
	errExternal = errors.New("repository: external service error")
)
