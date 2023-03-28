package auth

import (
	"drunklish/internal/config"
	"drunklish/internal/model"
	"drunklish/pkg/repository"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExistEmail(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		t.Log(err)
	}

	authService := NewAuthService(db)

	user, err := authService.SignIn(&model.User{
		Email:        "test@yandex.ru",
		HashPassword: "password",
	})
	if err != nil {
		t.Log(err)
	}

	_, err = authService.SignIn(&model.User{
		Email:        "test123@gmail.com",
		HashPassword: "password",
	})
	if err != nil {
		assert.ErrorIs(t, err, ErrEmailOrPassword)
	}

	err1 := authService.ExistEmail(db, user.user.Email)
	assert.NoError(t, err1)
}
