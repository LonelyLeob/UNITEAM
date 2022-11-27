package forms

import "errors"

var (
	errNotCorrectUUID     = errors.New("forms: entered UUID not in correct values, please enter correct")
	errCantShowForm       = errors.New("forms: changed form was deleted or patched, please change another form")
	errNoForms            = errors.New("forms: user have no forms, please create:)")
	errUnexpectedMethod   = errors.New("forms: unexpected signing method")
	errParseInt           = errors.New("forms: given number not may int")
	errAuthHeaderNotFound = errors.New("forms: cant found needed header in request")
	errAuthHeaderInvalid  = errors.New("forms: cant parse requested header")
	errCantParseToken     = errors.New("forms: cant parse token")
)
