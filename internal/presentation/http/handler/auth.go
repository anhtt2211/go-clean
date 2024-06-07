package handler

import (
	"My-Clean/internal/application/use-cases"
	"My-Clean/internal/presentation/http/inputs"
	"encoding/json"
	"net/http"

	"My-Clean/internal/utils"
)

type AuthHandler struct {
	AuthUseCase *use_cases.AuthUseCase
}

func NewAuthHandler(authUseCase *use_cases.AuthUseCase) *AuthHandler {
	return &AuthHandler{AuthUseCase: authUseCase}
}

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user inputs.RegisterInput
	json.NewDecoder(r.Body).Decode(&user)
	err := h.AuthUseCase.Register(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, user)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user inputs.LoginInput
	json.NewDecoder(r.Body).Decode(&user)
	token, err := h.AuthUseCase.Login(user.Username, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}
