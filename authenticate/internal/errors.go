package internal

import "errors"

var (
	errUnexpectedMethod = errors.New("cant parse with this method")
	errParseConflict    = errors.New("parse process stopped by unavailable reason")
	errInfo             = errors.New("cant show info")
)
