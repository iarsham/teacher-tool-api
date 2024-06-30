package handlers

import (
	"github.com/iarsham/bindme"
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"github.com/iarsham/teacher-tool-api/internal/entities"
	"github.com/iarsham/teacher-tool-api/internal/helpers"
	"net/http"
)

type RefreshHandler struct {
	Usecase domain.RefreshTokenUsecase
}

// RefreshTokenHandler godoc
//
//	@Summary		Refresh token
//	@Description	Refresh token endpoint to get new access token with refresh token
//	@Accept			json
//	@Produce		json
//	@Tags			Auth
//	@Param			tokenRequest	body		entities.RefreshTokenRequest	true	"Token data"
//	@Success		200				{object}	response.RefreshSuccess
//	@Failure		400				{object}	response.BadRequest
//	@Failure		500				{object}	response.InternalServerError
//	@router			/auth/refresh-token [post]
func (r *RefreshHandler) RefreshTokenHandler(w http.ResponseWriter, req *http.Request) {
	data := new(entities.RefreshTokenRequest)
	if err := bindme.ReadJson(req, data); err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	validateToken, err := r.Usecase.ValidateRefreshToken(data.RefreshToken)
	if err != nil {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": err.Error()}, nil)
		return
	}
	userID, ok := validateToken["sub"]
	if !ok {
		bindme.WriteJson(w, http.StatusBadRequest, helpers.M{"error": "invalid token. just refresh token allowed"}, nil)
		return
	}
	user, err := r.Usecase.FindById(uint64(userID.(float64)))
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	newToken, err := r.Usecase.CreateAccessToken(user.ID, user.Phone, user.Role)
	if err != nil {
		bindme.WriteJson(w, http.StatusInternalServerError, helpers.M{"error": helpers.ErrInternalServer.Error()}, nil)
		return
	}
	bindme.WriteJson(w, http.StatusOK, helpers.M{"access_token": newToken}, nil)
}
