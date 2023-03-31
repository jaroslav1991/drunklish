package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
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
)

func (a *Auth) SignIn(req model.User) (*dto.ResponseUser, error) {
	if existEmail := validator.ExistEmail(a.db, req.Email); existEmail == true {
		return nil, fmt.Errorf("%w", ErrEmail)
	}

	checkUser, err := a.repo.CheckUserDB(req)
	if err != nil {
		return nil, err
	}

	if checkPassword := users.CheckPasswordHash(checkUser.User.HashPassword, req.HashPassword); checkPassword != nil {
		return nil, fmt.Errorf("%w", ErrPassword)
	}

	newToken, err := users.GenerateToken(checkUser.User.Id, checkUser.User.Email)
	if err != nil {
		return nil, fmt.Errorf("%w", users.InvalidToken)
	}
	checkUser.Token = newToken

	return checkUser, nil
}
