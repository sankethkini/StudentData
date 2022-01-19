package memory

import "errors"

var (
	RecordNotFoundErr = errors.New("record not found")
	NotARightTypeErr  = errors.New("not a right type")
)
