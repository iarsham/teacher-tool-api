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

type RegisterHandler struct {
	Usecase domain.RegisterUsecase
}

// RegisterHandler godoc
//
//	@Summary		Register
//	@Description	Register a new user
//	@Accept			json
//	@Produce		json
//	@Tags			Auth
//	@Param			userRequest	body		entities.UserRequest	true	"User data"
//	@Success		201			{object}	response.UserCreated
//	@Failure		400			{object}	response.BadRequest
//	@Failure		409			{object}	response.UserAlreadyExists
//	@Failure		500			{object}	response.InternalServerError
//	@Router			/register [post]
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
