package application

import "errors"

var (
	NoNameErr     = errors.New("No proper name provided")
	AgeErr        = errors.New("Not a valid age")
	NoAddressErr  = errors.New("No proper address provided")
	NoRollNumErr  = errors.New("No proper rollnumber provided")
	RollExistsErr = errors.New("Roll number already exists")
)
