package exception

import "errors"

var (
	ErrUserNotExits        = errors.New("account is not found")
	ErrInternalServerError = errors.New("something went wrong")
)
