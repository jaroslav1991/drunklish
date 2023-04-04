package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
)

type SignInFn func(req model.User) (*dto.ResponseUser, error)

func SignInHandler(service AuthService) SignInFn {
	return func(req model.User) (*dto.ResponseUser, error) {
		authorizeUser, err := service.SignIn(req)
		if err != nil {
			return nil, err
		}

		return authorizeUser, nil
	}
}
