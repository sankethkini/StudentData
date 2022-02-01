package memory

import "errors"

var (
	ErrRecordNotFound = errors.New("record not found")
	ErrNotARightType  = errors.New("not a right type")
)
