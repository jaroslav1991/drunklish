package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth/dto"
	"fmt"
)

func (a *Auth) SignIn(req model.User) (*dto.ResponseUser, error) {
	existEmail, err := a.repo.ExistEmail(req.Email)
	if !existEmail {
		return nil, fmt.Errorf("invalid email: %w", httputils.ErrValidation)
	}
	if err != nil {
		return nil, httputils.ErrValidation
	}

	checkUser, err := a.repo.CheckUserDB(req)
	if err != nil {
		return nil, fmt.Errorf("invalid check user DB: %w", httputils.ErrValidation)
	}

	if checkPassword := a.checkPasswordFn(checkUser.User.HashPassword, req.HashPassword); checkPassword != nil {
		return nil, fmt.Errorf("invalid password hash: %w", httputils.ErrInternalServer)
	}

	newToken, err := a.generateTokenFn(checkUser.User.Id, checkUser.User.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid generate token: %w", httputils.ErrInternalServer)
	}
	checkUser.Token = newToken

	return checkUser, nil
}
