package repository

import "errors"

var (
	ErrorRecordNotFound = errors.New("record not found")
	ErrorRecordDeleted  = errors.New("record deleted")
)
