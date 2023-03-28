package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/users"
	"errors"
	"fmt"
)

const (
	authorizeQuery = `select * from users where email=$1`
)

var (
	ErrEmailOrPassword = errors.New("invalid email or password")
)

type ResponseUser struct {
	user  model.User
	Token string
}

func (a *Auth) SignIn(req *model.User) (*ResponseUser, error) {
	var passwordHash string
	var user ResponseUser

	if errEmail := a.ExistEmail(a.db, req.Email); errEmail != nil {
		return nil, fmt.Errorf("%w", ErrEmailOrPassword)
	}

	if err := a.db.QueryRowx(authorizeQuery, req.Email).Scan(&user.user.Id, &user.user.Email, &passwordHash); err != nil {
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
