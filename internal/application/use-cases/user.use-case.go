package use_cases

import (
	"My-Clean/internal/domain/entities"
)

type UserUseCase struct {
	UserRepository entities.UserRepository
}

func NewUserUseCase(repo entities.UserRepository) *UserUseCase {
	return &UserUseCase{UserRepository: repo}
}

func (uc *UserUseCase) GetByID(id int) (*entities.User, error) {
	return uc.UserRepository.GetByID(id)
}

func (uc *UserUseCase) GetByUsername(username string) (*entities.User, error) {
	return uc.UserRepository.GetByUsername(username)
}

func (uc *UserUseCase) GetAll() ([]*entities.User, error) {
	return uc.UserRepository.GetAll()
}

func (uc *UserUseCase) Create(user *entities.User) error {
	return uc.UserRepository.Create(user)
}

func (uc *UserUseCase) Update(user *entities.User) error {
	return uc.UserRepository.Update(user)
}

func (uc *UserUseCase) Delete(id int) error {
	return uc.UserRepository.Delete(id)
}
