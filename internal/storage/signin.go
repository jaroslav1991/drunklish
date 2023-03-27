package storage

import (
	"drunklish/internal/users"
	"errors"
)

const (
	authorizeQuery = `select * from users where email=$1`
)

type ResponseUser struct {
	user  User
	Token string
}

func (s *Storage) SignIn(req *User) (*ResponseUser, error) {
	var passwordHash string
	var user ResponseUser

	if err := s.DB.QueryRowx(authorizeQuery, req.Email).Scan(&user.user.Id, &user.user.Email, &passwordHash); err != nil {
		return nil, err
	}

	checkPassword := users.CheckPasswordHash(req.HashPassword, passwordHash)
	if !checkPassword {
		return nil, errors.New("invalid password")
	}

	newToken, err := users.GenerateToken(user.user.Id, user.user.Email)
	if err != nil {
		return nil, err
	}
	user.Token = newToken

	return &user, nil
}
