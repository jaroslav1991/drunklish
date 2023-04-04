package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
)

type AuthFn func(req dto.SignUpRequest) (*model.User, error)

func SignUpHandler(service AuthService) AuthFn {
	return func(req dto.SignUpRequest) (*model.User, error) {
		createdUser, err := service.SignUp(req)
		if err != nil {
			return nil, err
		}

		return createdUser, nil
	}
}
