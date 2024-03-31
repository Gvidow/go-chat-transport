package api

import "errors"

var (
	ErrHandlerNil = errors.New("handlers is nil")
	ErrServerNil  = errors.New("server is nil")
)
