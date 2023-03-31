package auth

import (
	"drunklish/internal/service"
	"drunklish/internal/service/auth/repository"
)

type Auth struct {
	db   service.DB
	repo *repository.AuthRepository
}

func NewAuthService(db service.DB, repo *repository.AuthRepository) *Auth {
	return &Auth{db: db, repo: repo}
}
