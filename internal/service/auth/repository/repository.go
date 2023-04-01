package repository

import (
	"database/sql"
	"drunklish/internal/model"
	"drunklish/internal/service"
	"drunklish/internal/service/auth/dto"
)

const (
	createUserQuery = `insert into users (email, hash_password) values ($1, $2) returning id`
	authorizeQuery  = `select * from users where email=$1`
	getEmail        = `select email from users where email=$1`
)

type AuthRepository struct {
	db service.DB
}

func NewAuthRepository(db service.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (repo *AuthRepository) CreateUser(userDTO dto.SignUpRequest) (*model.User, error) {
	var user model.User
	if err := repo.db.QueryRowx(createUserQuery, userDTO.Email, userDTO.Password).Scan(&user.Id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (repo *AuthRepository) CheckUserDB(user model.User) (*dto.ResponseUser, error) {
	if err := repo.db.QueryRowx(authorizeQuery, user.Email).Scan(&user.Id, &user.Email, &user.HashPassword); err != nil {
		return nil, err
	}

	return &dto.ResponseUser{User: user}, nil
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
