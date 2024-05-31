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

type UserHandler struct {
	Usecase domain.UserUsecase
}

func (u *UserHandler) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := helpers.GetUserID(r)
	user, err := u.Usecase.FindById(userID)
	switch {
	case errors.Is(sql.ErrNoRows, err):
		bindme.WriteJson(w, http.StatusNotFound, helpers.M{"error": "user not found"}, nil)
		return
	case err != nil:
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusOK, user, nil)
}

func (u *UserHandler) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := helpers.GetUserID(r)
	data := new(entities.UpdateUserRequest)
	if err := bindme.ReadJson(r, data); err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	user, err := u.Usecase.Update(userID, data)
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusOK, user, nil)
}

func (u *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := helpers.GetUserID(r)
	if err := u.Usecase.Delete(userID); err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusNoContent, nil, nil)
}
