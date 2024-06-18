package handler

import (
	"My-Clean/internal/application/use-cases"
	"My-Clean/internal/domain/entities"
	"My-Clean/internal/presentation/http/inputs"
	"My-Clean/internal/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserUseCase *use_cases.UserUseCase
}

func NewUserHandler(userUseCase *use_cases.UserUseCase) *UserHandler {
	return &UserHandler{UserUseCase: userUseCase}
}

func (h *UserHandler) GetMe(w http.ResponseWriter, r *http.Request) {
	// Get user from request context
	user, ok := r.Context().Value("user").(entities.User)
	if !ok {
		http.Error(w, "User not found in context", http.StatusInternalServerError)
		return
	}
	userID := user.ID

	// Retrieve user data from the database by ID
	userData, err := h.UserUseCase.GetByID(userID)
	if err != nil {
		http.Error(w, "Error retrieving user data", http.StatusInternalServerError)
		return
	}
	if userData == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Respond with user data
	json.NewEncoder(w).Encode(userData)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserUseCase.GetAll()
	if err != nil {
		http.Error(w, "Failed to retrieve users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if users == nil {
		utils.RespondWithJSON(w, http.StatusOK, []entities.User{}, "No users found")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, users, "Users retrieved successfully")
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch user by ID from use case
	user, err := h.UserUseCase.GetByID(id)
	if err != nil {
		http.Error(w, "Failed to retrieve user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if user was found
	if user == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// Respond with user data
	utils.RespondWithJSON(w, http.StatusOK, user, "User retrieved successfully")
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user inputs.CreateUserInput
	json.NewDecoder(r.Body).Decode(&user)
	err := h.UserUseCase.Create(user.ToCreateUserDto())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, user, "User created successfully")
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Decode the JSON request body into an UpdateUserInput
	var input inputs.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input data", http.StatusBadRequest)
		return
	}

	// Convert UpdateUserInput to UpdateUserDto
	userDto := input.ToUpdateUserDto()

	// Set the ID from the URL parameters
	userDto.ID = id

	// Call the use case to update the user
	err = h.UserUseCase.Update(userDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the updated user DTO
	utils.RespondWithJSON(w, http.StatusOK, userDto, "User updated successfully")
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}
	err = h.UserUseCase.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"}, "User deleted successfully")
}
