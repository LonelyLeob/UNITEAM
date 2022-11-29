package internal

import "errors"

var (
	errUnexpectedMethod   = errors.New("forms: unexpected signing method")
	errAuthHeaderNotFound = errors.New("forms: cant found needed header in request")
	errAuthHeaderInvalid  = errors.New("forms: cant parse requested header")
	errCantParseToken     = errors.New("forms: cant parse token")
)
