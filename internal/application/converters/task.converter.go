package converters

import "My-Clean/internal/domain/entities"

type TaskConverter interface {
	ToTaskEntity() *entities.Task
}
