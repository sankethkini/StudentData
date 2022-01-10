package application

import "errors"

var NoNameErr = errors.New("No proper name provided")
var AgeErr = errors.New("Not a valid age")
var NoAddress = errors.New("No proper address provided")
var NoRollNum = errors.New("No proper rollnumber provided")
var RollExists = errors.New("Roll number already exists")
