package dtos

import "My-Clean/internal/domain/entities"

type CreateUserDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserDto struct {
	ID       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (dto *CreateUserDto) ToUserEntity() *entities.User {
	return &entities.User{
		Username: dto.Username,
		Password: dto.Password,
	}
}
func (dto *UpdateUserDto) ToUserEntity() *entities.User {
	return &entities.User{
		ID:       dto.ID,
		Username: dto.Username,
		Password: dto.Password,
	}
}
