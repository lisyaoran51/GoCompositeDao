package goCompositeDao

import "errors"

var (
	ErrInternal       error = errors.New("internal error")
	ErrNotImplemented error = errors.New("not implemented error")
)
