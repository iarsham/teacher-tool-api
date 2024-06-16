package models

import "strconv"

type Grade uint64

const (
	Grade1 = iota
	Grade2
	Grade3
	Grade4
	Grade5
)

func (g Grade) String() string {
	switch g {
	case Grade1:
		return "first"
	case Grade2:
		return "second"
	case Grade3:
		return "third"
	case Grade4:
		return "fourth"
	case Grade5:
		return "fifth"
	default:
		return "unknown"
	}
}

func (g Grade) MarshalJSON() ([]byte, error) {
	quotedValue := strconv.Quote(g.String())
	return []byte(quotedValue), nil
}
