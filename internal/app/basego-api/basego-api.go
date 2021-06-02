package cstd

import "fmt"

// Defines app.
const (
	AppName = "basego-api"
	Version = "1.0.0"
)

// Defines app.
var (
	UserAgent = fmt.Sprintf("%s/%s", AppName, Version)
)
