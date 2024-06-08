package dtos

import "My-Clean/internal/domain/entities"

type RegisterDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto *RegisterDto) ToUserEntity() *entities.User {
	return &entities.User{
		Username: dto.Username,
		Password: dto.Password,
	}
}
