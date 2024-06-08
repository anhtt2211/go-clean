package use_cases

import (
	"My-Clean/internal/application/dtos"
	"My-Clean/internal/domain/entities"
	"My-Clean/internal/domain/repositories"
)

type UserUseCase struct {
	UserRepository repositories.UserRepository
}

func NewUserUseCase(repo repositories.UserRepository) *UserUseCase {
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

func (uc *UserUseCase) Create(user *dtos.CreateUserDto) error {
	return uc.UserRepository.Create(user.ToUserEntity())
}

func (uc *UserUseCase) Update(user *dtos.UpdateUserDto) error {
	return uc.UserRepository.Update(user.ToUserEntity())
}

func (uc *UserUseCase) Delete(id int) error {
	return uc.UserRepository.Delete(id)
}
