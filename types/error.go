package types

import "errors"

var (
	ErrArgs           = errors.New("args error")
	ErrInvalidRequest = errors.New("invalid request")
	ErrInvalidCommand = errors.New("invalid command")
	ErrExpiration     = errors.New("invalid expiration")
)
