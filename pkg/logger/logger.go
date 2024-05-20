package logger

import (
	"go.uber.org/zap"
)

func NewZapLog(debug bool) *zap.Logger {
	if debug {
		return zap.Must(zap.NewDevelopment())
	}
	return zap.Must(zap.NewProduction())
}
