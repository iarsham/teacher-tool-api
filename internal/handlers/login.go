package handlers

import (
	"database/sql"
	"errors"
	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/pkg/response"
	"go.uber.org/zap"
	"net/http"
)

type LoginHandler struct {
	Usecase domain.LoginUsecase
	Logger  *zap.Logger
}

func (a *LoginHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := new(entities.UserRequest)
	if err := bindme.ReadJson(r, data); err != nil {
		response.BadRequestJSON(w, a.Logger, err)
		return
	}
	user, err := a.Usecase.FindByPhone(data.Phone)
	if errors.Is(err, sql.ErrNoRows) {
		response.ErrJSON(w, http.StatusNotFound, a.Logger, "user not found")
		return
	}
	if err := a.Usecase.VerifyPass(user, data.Password); err != nil {
		response.ErrJSON(w, http.StatusUnauthorized, a.Logger, "wrong password")
		return
	}
	response.JSON(w, http.StatusOK, a.Logger, "user logged in", nil)
}
