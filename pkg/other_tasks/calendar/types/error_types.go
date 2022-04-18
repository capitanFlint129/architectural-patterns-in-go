package types

import "errors"

var (
	ErrorMethodNotAllowed = errors.New("method not allowed")
	ErrorEventNotFound    = errors.New("error not found")
)
