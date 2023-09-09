package httputils

import "errors"

var (
	ErrExistEmail     = errors.New("email is already exists")
	ErrInternalServer = errors.New("internal server error")
	ErrValidation     = errors.New("validation failed")
	UnmarshalError    = errors.New("unmarshal error")
	ReadBodyError     = errors.New("can't read body")
	ErrWordNotExist   = errors.New("word not exists")
	ParseMultiForm    = errors.New("parse multiform error")
	FormFile          = errors.New("formFile error")
	CreateFile        = errors.New("create file error")
	CopyFile          = errors.New("copy file error")
)
