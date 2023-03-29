package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/users"
	"drunklish/internal/service/auth/validator"
	"errors"
	"fmt"
)

const (
	authorizeQuery = `select * from users where email=$1`
)

var (
	ErrEmail    = errors.New("email is not exists")
	ErrPassword = errors.New("invalid password")
	//InvalidToken = errors.New("invalid token")
)

type ResponseUser struct {
	User  model.User
	Token string
}

func (a *Auth) SignIn(req *model.User) (*ResponseUser, error) {
	var passwordHash string
	var user ResponseUser

	if existEmail := validator.ExistEmail(a.db, req.Email); existEmail == true {
		return nil, fmt.Errorf("%w", ErrEmail)
	}

	if err := a.db.QueryRowx(authorizeQuery, req.Email).Scan(&user.User.Id, &user.User.Email, &passwordHash); err != nil {
		return nil, err
	}

	if checkPassword := users.CheckPasswordHash(req.HashPassword, passwordHash); checkPassword != nil {
		return nil, fmt.Errorf("%w", ErrPassword)
	}

	newToken, err := users.GenerateToken(user.User.Id, user.User.Email)
	if err != nil {
		return nil, fmt.Errorf("%w", users.InvalidToken)
	}
	user.Token = newToken

	return &user, nil
}
