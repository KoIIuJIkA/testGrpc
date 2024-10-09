package storage

import "errors"

var (
	ErrUserExist    = errors.New("user already exist")
	ErrUserNotFound = errors.New("user not found")
	ErrAppNotFFound = errors.New("app not found")
)
