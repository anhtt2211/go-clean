package use_cases

import (
	"My-Clean/internal/application/dtos"
	"My-Clean/internal/domain/entities"
	"My-Clean/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	UserRepository entities.UserRepository
}

func NewAuthUseCase(repo entities.UserRepository) *AuthUseCase {
	return &AuthUseCase{UserRepository: repo}
}

func (uc *AuthUseCase) Register(user *dtos.RegisterInput) error {
	//func (uc *AuthUseCase) Register(user *entities.User) error {
	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return uc.UserRepository.Create(user)
}

func (uc *AuthUseCase) Login(username, password string) (string, error) {
	user, err := uc.UserRepository.GetByUsername(username)
	if err != nil || user == nil {
		return "", err
	}
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return "", err
	}
	token, err := utils.GenerateJWT(*user)
	if err != nil {
		return "", err
	}
	return token, nil
}
