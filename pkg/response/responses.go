package response

import (
	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"go.uber.org/zap"
	"net/http"
)

const (
	Ok  string = "response"
	Err string = "error"
)

func ErrJSON(w http.ResponseWriter, status int, logger *zap.Logger, data interface{}) {
	if err := bindme.WriteJson(w, status, helpers.M{Err: data}, nil); err != nil {
		logger.Error(err.Error())
	}
}

func JSON(w http.ResponseWriter, status int, logger *zap.Logger, data interface{}, headers http.Header) {
	if err := bindme.WriteJson(w, status, helpers.M{Ok: data}, headers); err != nil {
		logger.Error(err.Error())
	}
}

func ServerErrJSON(w http.ResponseWriter, logger *zap.Logger, err error) {
	logger.Error(err.Error())
	msg := "the server encountered a problem and could not process your request."
	ErrJSON(w, http.StatusInternalServerError, logger, msg)
}

func BadRequestJSON(w http.ResponseWriter, logger *zap.Logger, err error) {
	ErrJSON(w, http.StatusBadRequest, logger, err.Error())
}
