package inputs

import "My-Clean/internal/application/dtos"

type CreateUserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUserInput struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (input *CreateUserInput) ToCreateUserDto() *dtos.CreateUserDto {
	return &dtos.CreateUserDto{
		Username: input.Username,
		Password: input.Password,
	}
}

func (input *UpdateUserInput) ToUpdateUserDto() *dtos.UpdateUserDto {
	return &dtos.UpdateUserDto{
		ID:       input.ID,
		Username: input.Username,
		Password: input.Password,
	}
}
