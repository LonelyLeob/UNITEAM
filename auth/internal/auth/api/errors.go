package api

import "errors"

var (
	ErrCantSignString   = errors.New("key is invalid")
	ErrTokenIsOut       = errors.New("token was set but ignored in map")
	ErrNoInsertMeta     = errors.New("meta is out")
	ErrNoResetMeta      = errors.New("cant reset on this meta")
	ErrUnexpectedMethod = errors.New("cant parse with this method")
	ErrParseConflict    = errors.New("parse process stopped by unavailable reason")
	ErrInfo             = errors.New("cant show info")
	ErrHeaderInvalid    = errors.New("header must be non-empty and be bearer, please enter the ticket")
)

type ErrorResponse struct {
	Message string `json:"err_message"`
}
