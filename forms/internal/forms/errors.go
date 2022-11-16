package forms

import "errors"

var (
	errNotCorrectUUID     = errors.New("entered UUID not in correct values, please enter correct")
	errCantShowForm       = errors.New("changed form was deleted or patched, please change another form")
	errNoForms            = errors.New("user have no forms, please create:)")
	errBadToParseToken    = errors.New("user give a bad token, please relog in auth service")
	errUnexpectedMethod   = errors.New("unexpected signing method")
	errParseInt           = errors.New("given number not may int")
	errAuthHeaderNotFound = errors.New("cant found needed header in request")
	errAuthHeaderInvalid  = errors.New("cant parse requested header")
)
