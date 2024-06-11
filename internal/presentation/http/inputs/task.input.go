package inputs

import (
	"My-Clean/internal/application/dtos"
	"My-Clean/internal/domain/entities"
	"github.com/go-playground/validator/v10"
	"time"
)

var validate = validator.New()

type CreateTaskInput struct {
	Title       string     `json:"title" binding:"required"`
	Description *string    `json:"description"`
	Priority    *string    `json:"priority" validate:"omitempty,oneof=Low Medium High"`
	DueDate     *time.Time `json:"due_date" validate:"omitempty"`
	Reminder    *time.Time `json:"reminder" validate:"omitempty"`
	Status      *string    `json:"status" validate:"omitempty,oneof=Pending Completed"`
	UserID      int        `json:"user_id" binding:"required"`
}

type UpdateTaskInput struct {
	ID          uint       `json:"id" binding:"required"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Priority    *string    `json:"priority" validate:"omitempty,oneof=Low Medium High"`
	DueDate     *time.Time `json:"due_date" validate:"omitempty"`
	Reminder    *time.Time `json:"reminder" validate:"omitempty"`
	Status      *string    `json:"status" validate:"omitempty,oneof=Pending Completed"`
}

func (input *CreateTaskInput) ToCreateTaskDto() (*dtos.CreateTaskDto, error) {
	var priority *entities.Priority
	if input.Priority != nil {
		p := entities.Priority(*input.Priority)
		priority = &p
	}

	var status *entities.Status
	if input.Status != nil {
		s := entities.Status(*input.Status)
		status = &s
	}

	// Validate the input
	if err := validate.Struct(input); err != nil {
		return nil, err
	}

	return &dtos.CreateTaskDto{
		Title:       input.Title,
		Description: input.Description,
		Priority:    priority,
		DueDate:     input.DueDate,
		Reminder:    input.Reminder,
		Status:      status,
		UserID:      input.UserID,
	}, nil
}

func (input *UpdateTaskInput) ToUpdateTaskDto() (*dtos.UpdateTaskDto, error) {
	var priority *entities.Priority
	if input.Priority != nil {
		p := entities.Priority(*input.Priority)
		priority = &p
	}

	var status *entities.Status
	if input.Status != nil {
		s := entities.Status(*input.Status)
		status = &s
	}

	// Validate the input
	if err := validate.Struct(input); err != nil {
		return nil, err
	}

	return &dtos.UpdateTaskDto{
		Title:       input.Title,
		Description: input.Description,
		Priority:    priority,
		DueDate:     input.DueDate,
		Reminder:    input.Reminder,
		Status:      status,
	}, nil
}
