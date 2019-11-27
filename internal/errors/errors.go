package errors

import "errors"

var ErrNotFound = errors.New("Not found")

var ErrConflict = errors.New("Conflict")

var ErrInvalidPayload = errors.New("Invalid payload")

var ErrInternalServer = errors.New("Internal Server")
