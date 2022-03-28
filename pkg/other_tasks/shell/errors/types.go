package errors

import "errors"

var (
	ErrorTooManyArguments   = errors.New("too many arguments")
	ErrorNotEnoughArguments = errors.New("not enough arguments")
)
