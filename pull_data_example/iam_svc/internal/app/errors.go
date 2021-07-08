package app

import "errors"

// Common and generic errors.
var (
	ErrDuplicateEmail    = errors.New("duplicate email")
	ErrRecordNotFound    = errors.New("record not found")
	ErrSubjectEIDInvalid = errors.New("subject's provided id is invalid")
)
