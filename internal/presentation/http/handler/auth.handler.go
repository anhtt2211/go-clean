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
	var userInput inputs.RegisterInput
	if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.AuthUseCase.Register(userInput.ToRegisterDto()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, map[string]string{"message": "User registered successfully"})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginInput inputs.LoginInput
	if err := json.NewDecoder(r.Body).Decode(&loginInput); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	token, err := h.AuthUseCase.Login(loginInput.ToLoginDto())
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}
