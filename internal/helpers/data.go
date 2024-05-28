package helpers

import "errors"

type M map[string]interface{}

var ErrInternalServer = errors.New("the server encountered a problem and could not process your request")
