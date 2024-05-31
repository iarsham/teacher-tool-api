package handlers

import (
	"database/sql"
	"errors"
	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"net/http"
)

type LoginHandler struct {
	Usecase domain.LoginUsecase
}

func (a *LoginHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	data := new(entities.UserRequest)
	if err := bindme.ReadJson(r, data); err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	user, err := a.Usecase.FindByPhone(data.Phone)
	if errors.Is(err, sql.ErrNoRows) {
		bindme.WriteJson(w, http.StatusNotFound, helpers.M{"error": "user not found"}, nil)
		return
	}
	if err := a.Usecase.VerifyPass(user.Password, data.Password); err != nil {
		bindme.WriteJson(w, http.StatusUnprocessableEntity, helpers.M{"error": "wrong password"}, nil)
		return
	}
	accessToken, err := a.Usecase.CreateAccessToken(user.ID, user.Phone)
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	refreshToken, err := a.Usecase.CreateRefreshToken(user.ID)
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	tokens := helpers.M{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}
	bindme.WriteJson(w, http.StatusOK, tokens, nil)
}
