package errors

import "errors"

// TODO вопрос: разделять ли ошибки по пакетам? Или оставить общий файл на весь shell?

var (
	ErrorTooManyArguments   = errors.New("too many arguments")
	ErrorNotEnoughArguments = errors.New("not enough arguments")
	ErrorCommandNotFound    = errors.New("command not found")
)
