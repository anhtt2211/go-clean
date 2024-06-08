package converters

import "My-Clean/internal/domain/entities"

type UserConverter interface {
	ToUserEntity() *entities.User
}
