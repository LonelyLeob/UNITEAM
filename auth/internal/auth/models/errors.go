package models

import "errors"

var (
	errValidation    = errors.New("models: validation error")
	errWrongPassword = errors.New("models: wrong compare hash password")
	errNoUserAgent   = errors.New("models: wrong user agent")
)
