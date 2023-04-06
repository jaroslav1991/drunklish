package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
)

type SignInFn func(req model.User) (*dto.ResponseFromSignIn, error)

func SignInHandler(service AuthService) SignInFn {
	return func(req model.User) (*dto.ResponseFromSignIn, error) {
		authorizeUser, err := service.SignIn(req)
		if err != nil {
			return nil, err
		}

		return authorizeUser, nil
	}
}
