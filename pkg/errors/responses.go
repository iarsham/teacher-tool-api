package errors

import (
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"go.uber.org/zap"
	"net/http"
)

func errResponse(w http.ResponseWriter, status int, l *zap.Logger, data interface{}) {
	if err := helpers.WriteJson(w, status, helpers.M{"error": data}); err != nil {
		l.Error(err.Error())
	}
}

func ServerErrResponse(w http.ResponseWriter, l *zap.Logger, err error) {
	l.Error(err.Error())
	errResponse(w, http.StatusInternalServerError, l, http.StatusText(http.StatusInternalServerError))
}
