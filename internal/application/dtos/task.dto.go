package dtos

import (
	"My-Clean/internal/domain/entities"
	"time"
)

type CreateTaskDto struct {
	Title       string             `json:"title" binding:"required"`
	Description *string            `json:"description"`
	Priority    *entities.Priority `json:"priority"`
	DueDate     *time.Time         `json:"due_date"`
	Reminder    *time.Time         `json:"reminder"`
	Status      *entities.Status   `json:"status"`
	UserID      int                `json:"user_id" binding:"required"`
}

func (dto *CreateTaskDto) ToTaskEntity() *entities.Task {
	status := entities.Pending
	if dto.Status != nil {
		status = *dto.Status
	}

	return &entities.Task{
		Title:       dto.Title,
		Description: dto.Description,
		Priority:    dto.Priority,
		DueDate:     dto.DueDate,
		Reminder:    dto.Reminder,
		Status:      status,
		UserID:      dto.UserID,
	}
}

type UpdateTaskDto struct {
	ID          uint               `json:"id" binding:"required"`
	Title       string             `json:"title"`
	Description *string            `json:"description"`
	Priority    *entities.Priority `json:"priority"`
	DueDate     *time.Time         `json:"due_date"`
	Reminder    *time.Time         `json:"reminder"`
	Status      *entities.Status   `json:"status"`
}

func (dto *UpdateTaskDto) ToTaskEntity() *entities.Task {
	status := entities.Pending
	if dto.Status != nil {
		status = *dto.Status
	}

	return &entities.Task{
		Title:       dto.Title,
		Description: dto.Description,
		Priority:    dto.Priority,
		DueDate:     dto.DueDate,
		Reminder:    dto.Reminder,
		Status:      status,
	}
}
