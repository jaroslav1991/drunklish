package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth/dto"
	"fmt"
)

func (a *Auth) SignIn(req model.User) (*dto.ResponseFromSignIn, error) {
	existEmail, err := a.repo.ExistEmail(req.Email)
	if !existEmail {
		return nil, fmt.Errorf("email not exists: %w", httputils.ErrValidation)
	}
	if err != nil {
		return nil, httputils.ErrValidation
	}

	checkUser, err := a.repo.CheckUserDB(req)
	if err != nil {
		return nil, fmt.Errorf("invalid check user DB: %w", httputils.ErrValidation)
	}

	if err := a.checkPasswordFn(checkUser.User.HashPassword, req.HashPassword); err != nil {
		return nil, fmt.Errorf("invalid password: %w", httputils.ErrValidation)
	}

	newToken, err := a.generateTokenFn(checkUser.User.Id, checkUser.User.Email)
	if err != nil {
		return nil, fmt.Errorf("invalid generate token: %w", httputils.ErrInternalServer)
	}
	//checkUser.Token = newToken

	return &dto.ResponseFromSignIn{Token: newToken}, nil
}
