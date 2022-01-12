package adapters

import "errors"

var NoAdapterFound = errors.New("no matching adapters found")
var NotARightType = errors.New("not a right type")
var RecordNotFound = errors.New("record not found")
