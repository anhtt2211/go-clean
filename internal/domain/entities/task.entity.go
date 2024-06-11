package entities

import (
	"gorm.io/gorm"
	"time"
)

type Priority string

const (
	Low    Priority = "Low"
	Medium Priority = "Medium"
	High   Priority = "High"
)

type Status string

const (
	Pending   Status = "Pending"
	Completed Status = "Completed"
)

// Task struct with enums for priority and status
type Task struct {
	ID          int            `json:"id" gorm:"primaryKey"`
	Title       string         `json:"title"`
	Description *string        `json:"description"`                   // Nullable
	Priority    *Priority      `json:"priority"`                      // Nullable
	DueDate     *time.Time     `json:"due_date"`                      // Nullable
	Reminder    *time.Time     `json:"reminder"`                      // Nullable
	Status      Status         `json:"status" gorm:"default:Pending"` // Default value
	UserID      int            `json:"user_id"`
	CreatedAt   time.Time      `json:"created_at" gorm:"autoCreateTime"` // Automatically set `CreatedAt`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"autoUpdateTime"` // Automatically set `UpdatedAt`
	DeletedAt   gorm.DeletedAt `gorm:"index"`                            // Soft delete support
}
