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

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
	Logger      *zap.Logger
}

func (a *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := new(entities.UserRequest)
	if err := bindme.ReadJson(r, data); err != nil {
		response.BadRequestJSON(w, a.Logger, err)
		return
	}

	if _, err := a.AuthUsecase.FindByPhone(data.Phone); !errors.Is(err, sql.ErrNoRows) {
		response.ErrJSON(w, http.StatusConflict, a.Logger, "user already exists")
		return
	}

	hashPass, err := a.AuthUsecase.EncryptPass(data.Password)
	if err != nil {
		response.ServerErrJSON(w, a.Logger, err)
		return
	}
	data.Password = hashPass

	user, err := a.AuthUsecase.Create(data)
	if err != nil {
		response.ServerErrJSON(w, a.Logger, err)
		return
	}

	response.JSON(w, http.StatusCreated, a.Logger, user, nil)
}
