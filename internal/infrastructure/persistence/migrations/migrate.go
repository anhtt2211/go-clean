package migrations

import (
	"My-Clean/internal/domain/entities"
	"My-Clean/internal/infrastructure/persistence"
)

func AutoMigrate() error {
	err := persistence.DB.AutoMigrate(&entities.User{})
	if err != nil {
		return err
	}
	return nil
}
