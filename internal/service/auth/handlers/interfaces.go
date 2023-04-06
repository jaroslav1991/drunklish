package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
)

type AuthService interface {
	SignUp(req dto.SignUpRequest) (*model.User, error)
	SignIn(req model.User) (*dto.ResponseFromSignIn, error)
}
