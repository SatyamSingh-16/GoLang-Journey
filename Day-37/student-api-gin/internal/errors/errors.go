package errors

import "errors"

var ErrStudentNotFound = errors.New("student not found")
var ErrEmailAlreadyExists = errors.New("email already exists")

var ErrInvalidCredentials = errors.New("invalid credentials")
