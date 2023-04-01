package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
	"drunklish/internal/service/auth/token"
	"errors"
	"fmt"
)

var (
	ErrEmail    = errors.New("email is not exists")
	ErrPassword = errors.New("invalid password")
)

func (a *Auth) SignIn(req model.User) (*dto.ResponseUser, error) {
	existEmail, err := a.repo.ExistEmail(req.Email)
	if !existEmail {
		return nil, ErrEmail
	}
	if err != nil {
		return nil, err
	}

	checkUser, err := a.repo.CheckUserDB(req)
	if err != nil {
		return nil, fmt.Errorf("%w", ErrEmail)
	}

	if checkPassword := token.CheckPasswordHash(checkUser.User.HashPassword, req.HashPassword); checkPassword != nil {
		return nil, fmt.Errorf("%w", ErrPassword)
	}

	newToken, err := token.GenerateToken(checkUser.User.Id, checkUser.User.Email)
	if err != nil {
		return nil, fmt.Errorf("%w", token.InvalidToken)
	}
	checkUser.Token = newToken

	return checkUser, nil
}
