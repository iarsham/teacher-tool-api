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

// GetUserHandler godoc
//
//	@Summary		Get User
//	@Description	Get a user data and properties
//	@Accept			json
//	@Produce		json
//	@Tags			Users
//	@Success		200	{object}	response.UserData
//	@Failure		404	{object}	response.UserNotFound
//	@Failure		500	{object}	response.InternalServerError
//	@router			/user [get]
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

// UpdateUserHandler godoc
//
//	@Summary		Update User
//	@Description	Update a user based on jwt and current user
//	@Accept			json
//	@Produce		json
//	@Tags			Users
//	@Param			userRequest	body		entities.UpdateUserRequest	true	"User data"
//	@Success		200			{object}	response.UserData
//	@Failure		400			{object}	response.BadRequest
//	@Failure		500			{object}	response.InternalServerError
//	@router			/user [put]
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

// DeleteUserHandler godoc
//
//	@Summary		Delete User
//	@Description	Delete a user based on jwt and current user
//	@Accept			json
//	@Produce		json
//	@Tags			Users
//	@Success		204	{object}	nil
//	@Failure		500	{object}	response.InternalServerError
//	@router			/user [delete]
func (u *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	userID := helpers.GetUserID(r)
	if err := u.Usecase.Delete(userID); err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusNoContent, nil, nil)
}
