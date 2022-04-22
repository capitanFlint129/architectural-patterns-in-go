package types

import "errors"

var (
	ErrorMethodNotAllowed   = errors.New("method not allowed")
	ErrorEventNotFound      = errors.New("event not found")
	ErrorEventAlreadyExists = errors.New("event already exists")
)
