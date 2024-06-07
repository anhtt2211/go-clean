package repositories

import (
	"My-Clean/internal/domain/entities"
	"My-Clean/internal/domain/repositories"
	"gorm.io/gorm"
)

type GORMUserRepository struct {
	DB *gorm.DB
}

func NewGORMUserRepository(db *gorm.DB) repositories.UserRepository {
	return &GORMUserRepository{DB: db}
}

func (repo *GORMUserRepository) Create(user *entities.User) error {
	result := repo.DB.Create(user)
	return result.Error
}

func (repo *GORMUserRepository) GetByID(id int) (*entities.User, error) {
	var user entities.User
	result := repo.DB.First(&user, id)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, result.Error
}

func (repo *GORMUserRepository) GetByUsername(username string) (*entities.User, error) {
	var user entities.User
	result := repo.DB.Where("username = ?", username).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, result.Error
}

func (repo *GORMUserRepository) GetAll() ([]*entities.User, error) {
	var users []*entities.User
	result := repo.DB.Find(&users)
	return users, result.Error
}

func (repo *GORMUserRepository) Update(user *entities.User) error {
	result := repo.DB.Save(user)
	return result.Error
}

func (repo *GORMUserRepository) Delete(id int) error {
	result := repo.DB.Delete(&entities.User{}, id)
	return result.Error
}
