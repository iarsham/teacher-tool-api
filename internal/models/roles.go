package models

import (
	"strconv"
)

type Role uint64

const (
	teacher = iota
	student
	admin
)

func (r Role) MarshalJSON() ([]byte, error) {
	quotedValue := strconv.Quote(r.String())
	return []byte(quotedValue), nil
}

func (r Role) String() string {
	switch r {
	case teacher:
		return "teacher"
	case student:
		return "student"
	case admin:
		return "admin"
	default:
		return "unknown"
	}
}
