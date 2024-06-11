package use_cases

import (
	"My-Clean/internal/application/dtos"
	"My-Clean/internal/domain/entities"
	"My-Clean/internal/domain/repositories"
	"My-Clean/internal/domain/types"
)

type TaskUseCase struct {
	TaskRepository repositories.TaskRepository
}

func NewTaskUseCase(repo repositories.TaskRepository) *TaskUseCase {
	return &TaskUseCase{TaskRepository: repo}
}

func (uc *TaskUseCase) Gets(options types.PaginateOptions) ([]*entities.Task, error) {
	return uc.TaskRepository.Gets(options)
}

func (uc *TaskUseCase) GetByID(id uint) (*entities.Task, error) {
	return uc.TaskRepository.GetByID(id)
}

func (uc *TaskUseCase) Create(Task *dtos.CreateTaskDto) error {
	return uc.TaskRepository.Create(Task.ToTaskEntity())
}

func (uc *TaskUseCase) Update(Task *dtos.UpdateTaskDto) error {
	return uc.TaskRepository.Update(Task.ToTaskEntity())
}

func (uc *TaskUseCase) Delete(id uint) error {
	return uc.TaskRepository.Delete(id)
}
