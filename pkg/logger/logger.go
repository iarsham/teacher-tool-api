package logger

import "go.uber.org/zap"

func NewZapLog() *zap.Logger {
	return zap.Must(zap.NewDevelopment())
}
