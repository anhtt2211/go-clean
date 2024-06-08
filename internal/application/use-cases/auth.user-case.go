package use_cases

import (
	"My-Clean/internal/application/dtos"
	"My-Clean/internal/domain/repositories"
	"My-Clean/internal/utils"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase struct {
	UserRepository repositories.UserRepository
}

func NewAuthUseCase(repo repositories.UserRepository) *AuthUseCase {
	return &AuthUseCase{UserRepository: repo}
}

// Register handles user registration by hashing the password and saving the user.
func (uc *AuthUseCase) Register(userDto *dtos.RegisterDto) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("failed to hash password: " + err.Error())
	}

	userDto.Password = string(hashedPassword)

	if err := uc.UserRepository.Create(userDto.ToUserEntity()); err != nil {
		return errors.New("failed to create user: " + err.Error())
	}

	return nil
}

// Login handles user login by verifying the password and generating a JWT token.
func (uc *AuthUseCase) Login(loginDto *dtos.LoginDto) (string, error) {
	user, err := uc.UserRepository.GetByUsername(loginDto.Username)
	if err != nil {
		return "", errors.New("failed to fetch user: " + err.Error())
	}
	if user == nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginDto.Password)); err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(*user)
	if err != nil {
		return "", errors.New("failed to generate token: " + err.Error())
	}

	return token, nil
}
