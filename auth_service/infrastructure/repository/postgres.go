package repository

import (
	"chat_app/auth_service/domain"
	"database/sql"
	"errors"
)

type PostgresUserRepo struct {
	DB *sql.DB
}

func NewPostgresUserRepo(db *sql.DB) *PostgresUserRepo {
	return &PostgresUserRepo{DB: db}
}

func (r *PostgresUserRepo) CreateUser(user *domain.User) error {
	var exists bool
	err := r.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)", user.Username).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("пользователь уже существует")
	}

	query := "INSERT INTO users (username, password) VALUES ($1, $2)"
	_, err = r.DB.Exec(query, user.Username, user.Password)
	return err
}

func (r *PostgresUserRepo) GetUserByUsername(username string) (*domain.User, error) {
	user := &domain.User{}
	query := "SELECT id, username, password FROM users WHERE username = $1"
	row := r.DB.QueryRow(query, username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("пользователь не найден")
		}
		return nil, err
	}
	return user, nil
}
