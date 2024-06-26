package inputs

import "My-Clean/internal/application/dtos"

type RegisterInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (input *RegisterInput) ToRegisterDto() *dtos.RegisterDto {
	return &dtos.RegisterDto{
		Username: input.Username,
		Password: input.Password,
	}
}

func (input *LoginInput) ToLoginDto() *dtos.LoginDto {
	return &dtos.LoginDto{
		Username: input.Username,
		Password: input.Password,
	}
}
