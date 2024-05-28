package handlers

import (
	"database/sql"
	"errors"
	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"go.uber.org/zap"
	"net/http"
)

type RegisterHandler struct {
	Usecase domain.RegisterUsecase
	Logger  *zap.Logger
}

func (a *RegisterHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	data := new(entities.UserRequest)
	if err := bindme.ReadJson(r, data); err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	if _, err := a.Usecase.FindByPhone(data.Phone); !errors.Is(err, sql.ErrNoRows) {
		bindme.WriteJson(w, http.StatusConflict, helpers.M{"error": "user already exists"}, nil)
		return
	}
	hashPass, err := a.Usecase.EncryptPass(data.Password)
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	data.Password = string(hashPass)
	if _, err := a.Usecase.Create(data); err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusCreated, helpers.M{"response": "user created"}, nil)
}
