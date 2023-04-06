package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSignInHandler_Positive(t *testing.T) {
	service := &mockService{fnA: func(req model.User) (*dto.ResponseFromSignIn, error) {
		assert.Equal(t, int64(1), req.Id)
		assert.Equal(t, "qwerty", req.Email)
		assert.Equal(t, "123", req.HashPassword)

		return &dto.ResponseFromSignIn{
			Token: "token",
		}, nil
	}}

	expectedResponse := &dto.ResponseFromSignIn{
		Token: "token",
	}
	handler := SignInHandler(service)

	actualResponse, actualErr := handler(model.User{
		Id:           1,
		Email:        "qwerty",
		HashPassword: "123",
	})

	assert.NoError(t, actualErr)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestSignInHandler_Negative(t *testing.T) {
	service := &mockService{fnA: func(req model.User) (*dto.ResponseFromSignIn, error) {
		assert.Equal(t, int64(1), req.Id)
		assert.Equal(t, "qwerty", req.Email)
		assert.Equal(t, "123", req.HashPassword)

		return nil, errors.New("fuck up")
	}}

	handler := SignInHandler(service)

	_, actualErr := handler(model.User{
		Id:           1,
		Email:        "qwerty",
		HashPassword: "123",
	})

	assert.Error(t, actualErr)
}
