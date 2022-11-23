package internal

import "errors"

var (
	errUnexpectedMethod     = errors.New("cant parse with this method")
	errParseConflict        = errors.New("parse process stopped by unavailable reason")
	errInfo                 = errors.New("cant show info")
	errFieldsMustBeNotEmpty = errors.New("one of fields empty, please enter data")
	errEmptyHeader          = errors.New("header must be non-empty, please enter the ticket")
)
