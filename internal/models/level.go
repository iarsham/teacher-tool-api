package models

import "strconv"

type Level uint64

const (
	Easy = iota
	Medium
	Advanced
)

func (l Level) String() string {
	switch l {
	case Easy:
		return "easy"
	case Medium:
		return "medium"
	case Advanced:
		return "advanced"
	default:
		return "unknown"
	}
}

func (l Level) MarshalJSON() ([]byte, error) {
	quotedValue := strconv.Quote(l.String())
	return []byte(quotedValue), nil
}
