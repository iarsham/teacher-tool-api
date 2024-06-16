package models

import (
	"strconv"
)

type Role uint64

const (
	Teacher = iota
	Student
	Admin
)

func (r Role) String() string {
	switch r {
	case Teacher:
		return "teacher"
	case Student:
		return "student"
	case Admin:
		return "admin"
	default:
		return "unknown"
	}
}

func (r Role) MarshalJSON() ([]byte, error) {
	quotedValue := strconv.Quote(r.String())
	return []byte(quotedValue), nil
}
