package httputils

import "errors"

var (
	ErrExistEmail     = errors.New("email is already exists")
	ErrInternalServer = errors.New("internal server error")
	ErrValidation     = errors.New("validation failed")
	UnmarshalError    = errors.New("unmarshal error")
	ReadBodyError     = errors.New("can't read body")
	ErrWordNotExist   = errors.New("words not exists")
)
