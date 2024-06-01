package handlers

import (
	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"net/http"
)

type PasswordHandler struct {
	Usecase domain.PassUsecase
}

// PasswordChangeHandler godoc
//
//	@Summary		Password Change
//	@Description	Change user password
//	@Accept			json
//	@Produce		json
//	@Tags			Users
//	@Param			userRequest	body		entities.PassChangeRequest	true	"User data"
//	@Success		200			{object}	response.PasswordChanged
//	@Failure		400			{object}	response.BadRequest
//	@Failure		500			{object}	response.InternalServerError
//	@router			/user/change-password [post]
func (p *PasswordHandler) PasswordChangeHandler(w http.ResponseWriter, r *http.Request) {
	userID := helpers.GetUserID(r)
	data := new(entities.PassChangeRequest)
	if err := bindme.ReadJson(r, data); err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	if ok := data.PasswordsIsEqual(); !ok {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": "passwords is not equal"}, nil)
		return
	}
	hashPass, err := p.Usecase.EncryptPass(data.Password)
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	if _, err := p.Usecase.UpdatePassword(userID, string(hashPass)); err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusOK, helpers.M{"response": "password changed successfully"}, nil)
}
