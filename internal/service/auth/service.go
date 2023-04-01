package auth

import (
	"drunklish/internal/service/auth/repository"
)

type Auth struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo *repository.AuthRepository) *Auth {
	return &Auth{repo: repo}
}
