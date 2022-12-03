package api

import "errors"

var (
	errCantSignString   = errors.New("key is invalid")
	errTokenIsOut       = errors.New("token was set but ignored in map")
	errNoInsertMeta     = errors.New("meta is out")
	errUnexpectedMethod = errors.New("cant parse with this method")
	errParseConflict    = errors.New("parse process stopped by unavailable reason")
	errInfo             = errors.New("cant show info")
	errHeaderInvalid    = errors.New("header must be non-empty and be bearer, please enter the ticket")
)

type errorResponse struct {
	Message string `json:"err_message"`
}
