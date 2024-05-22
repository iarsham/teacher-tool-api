package handlers

import (
	"github.com/iarsham/teacher-tool-api/internal/domain"
	"net/http"
)

type AuthHandler struct {
	AuthUsecase domain.AuthUsecase
}

func (a *AuthHandler) LoginRegisterHandler(w http.ResponseWriter, r *http.Request) {

}
