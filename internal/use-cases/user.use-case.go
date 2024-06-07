package use_cases

import (
	"My-Clean/internal/domain"
	"My-Clean/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	UserRepository domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{UserRepository: repo}
}

func (uc *UserUseCase) Register(user *domain.User) error {
	// Hash the password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return uc.UserRepository.Create(user)
}

func (uc *UserUseCase) Login(username, password string) (string, error) {
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

// Other use case methods for CRUD operations

func (uc *UserUseCase) GetByID(id int) (*domain.User, error) {
	return uc.UserRepository.GetByID(id)
}

func (uc *UserUseCase) GetByUsername(username string) (*domain.User, error) {
	return uc.UserRepository.GetByUsername(username)
}

func (uc *UserUseCase) GetAll() ([]*domain.User, error) {
	return uc.UserRepository.GetAll()
}

func (uc *UserUseCase) Update(user *domain.User) error {
	return uc.UserRepository.Update(user)
}

func (uc *UserUseCase) Delete(id int) error {
	return uc.UserRepository.Delete(id)
}
