package repositories

import (
	"My-Clean/internal/domain/entities"
	"My-Clean/internal/domain/types"
)

type TaskRepository interface {
	Create(task *entities.Task) error
	GetByID(id uint) (*entities.Task, error)
	Gets(options types.PaginateOptions) ([]*entities.Task, error)
	Update(task *entities.Task) error
	Delete(id uint) error
}
