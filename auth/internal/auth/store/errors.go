package store

import "errors"

var (
	errNilPtr            = errors.New("store: repo get nil pointer struct")
	errUnreachableAction = errors.New("store: cant do action about user record")
)
