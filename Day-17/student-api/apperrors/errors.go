package errors

import "errors"

var ErrStudentNotFound = errors.New(
	"student not found",
)
var ErrStudentAlreadyExists = errors.New(
	"student already exists",
)
var ErrInvalidCredentials = errors.New(
	"invalid credentials",
)
