package repositories

import (
	"My-Clean/internal/domain/entities"
	"My-Clean/internal/domain/repositories"
	"My-Clean/internal/domain/types"
	"gorm.io/gorm"
	"time"
)

type GORMTaskRepository struct {
	DB *gorm.DB
}

func NewGORMTaskRepository(db *gorm.DB) repositories.TaskRepository {
	return &GORMTaskRepository{DB: db}
}

func (repo *GORMTaskRepository) Create(task *entities.Task) error {
	result := repo.DB.Create(task)
	return result.Error
}

func (repo *GORMTaskRepository) GetByID(id uint) (*entities.Task, error) {
	var task entities.Task
	result := repo.DB.First(&task, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &task, result.Error
}

func (repo *GORMTaskRepository) Gets(options types.PaginateOptions) ([]*entities.Task, error) {
	var tasks []*entities.Task
	query := repo.DB

	if options.Keyword != "" {
		query = query.Where("title LIKE ? OR description LIKE ?", "%"+options.Keyword+"%", "%"+options.Keyword+"%")
	}
	if options.Filters != nil {
		query = query.Where(options.Filters)
	}
	if options.OrderBy != "" {
		query = query.Order(options.OrderBy + " " + options.Order)
	}
	if options.Limit > 0 {
		query = query.Limit(options.Limit)
	}
	if options.Page > 0 {
		query = query.Offset((options.Page - 1) * options.Limit)
	}

	result := query.Find(&tasks)
	return tasks, result.Error
}

func (repo *GORMTaskRepository) Update(task *entities.Task) error {
	result := repo.DB.Save(task)
	return result.Error
}

func (repo *GORMTaskRepository) Delete(id uint) error {
	result := repo.DB.Delete(&entities.Task{}, id)
	return result.Error
}

func (repo *GORMTaskRepository) UpdateReminder(id uint, reminder time.Time) (bool, error) {
	result := repo.DB.Update("reminder", reminder).Where("id = ?", id)
	return result.RowsAffected > 0, result.Error
}
