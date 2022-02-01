package application

import "errors"

var (
	ErrNoName     = errors.New("no proper name provided")
	ErrAge        = errors.New("not a valid age")
	ErrNoAddress  = errors.New("no proper address provided")
	ErrNoRollNum  = errors.New("no proper rollnumber provided")
	ErrRollExists = errors.New("roll number already exists")
)
