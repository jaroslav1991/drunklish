package repository

import (
	"database/sql"
	"drunklish/internal/model"
	"drunklish/internal/pkg/db"
	"drunklish/internal/service/auth/dto"
)

const (
	createUserQuery = `insert into users (email, hash_password) values ($1, $2) returning id`
	authorizeQuery  = `select * from users where email=$1`
	getEmail        = `select email from users where email=$1`
)

type AuthRepository struct {
	db db.DB
}

func NewAuthRepository(db db.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (repo *AuthRepository) CreateUser(email, password string) (*model.User, error) {
	var user model.User
	if err := repo.db.QueryRowx(createUserQuery, email, password).Scan(&user.Id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *AuthRepository) CheckUserDB(email string) (*dto.ResponseUser, error) {
	var user dto.ResponseUser

	if err := repo.db.QueryRowx(authorizeQuery, email).Scan(&user.User.Id, &user.User.Email, &user.User.HashPassword); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *AuthRepository) ExistEmail(email string) (bool, error) {
	err := repo.db.QueryRowx(getEmail, email).Scan(&email)
	if err != nil && err == sql.ErrNoRows {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}
