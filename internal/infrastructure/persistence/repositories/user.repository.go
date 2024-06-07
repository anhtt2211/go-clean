package repositories

import (
	"My-Clean/internal/domain"
	"database/sql"
)

type MySQLUserRepository struct {
	DB *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) domain.UserRepository {
	return &MySQLUserRepository{DB: db}
}

func (repo *MySQLUserRepository) Create(user *domain.User) error {
	query := "INSERT INTO users (username, password) VALUES (?, ?)"
	_, err := repo.DB.Exec(query, user.Username, user.Password)
	return err
}

func (repo *MySQLUserRepository) GetByID(id int) (*domain.User, error) {
	query := "SELECT id, username, password FROM users WHERE id = ?"
	row := repo.DB.QueryRow(query, id)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (repo *MySQLUserRepository) GetByUsername(username string) (*domain.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	row := repo.DB.QueryRow(query, username)
	user := &domain.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return user, err
}

func (repo *MySQLUserRepository) GetAll() ([]*domain.User, error) {
	query := "SELECT id, username, password FROM users"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*domain.User{}
	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (repo *MySQLUserRepository) Update(user *domain.User) error {
	query := "UPDATE users SET username = ?, password = ? WHERE id = ?"
	_, err := repo.DB.Exec(query, user.Username, user.Password, user.ID)
	return err
}

func (repo *MySQLUserRepository) Delete(id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := repo.DB.Exec(query, id)
	return err
}
