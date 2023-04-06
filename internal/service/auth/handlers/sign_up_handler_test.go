package handlers

import (
	"drunklish/internal/model"
	"drunklish/internal/service/auth/dto"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

type mockService struct {
	fnR func(req dto.SignUpRequest) (*model.User, error)
	fnA func(req model.User) (*dto.ResponseFromSignIn, error)
}

func (m *mockService) SignUp(req dto.SignUpRequest) (*model.User, error) {
	return m.fnR(req)
}

func (m *mockService) SignIn(req model.User) (*dto.ResponseFromSignIn, error) {
	return m.fnA(req)
}

func TestSignUpHandler_Positive(t *testing.T) {
	service := &mockService{fnR: func(req dto.SignUpRequest) (*model.User, error) {
		assert.Equal(t, "qwerty", req.Email)
		assert.Equal(t, "123", req.Password)

		return &model.User{
			Id:           1,
			Email:        "qwerty",
			HashPassword: "123",
		}, nil
	}}

	expectedResponse := &model.User{
		Id:           1,
		Email:        "qwerty",
		HashPassword: "123",
	}

	handler := SignUpHandler(service)

	actualResponse, actualErr := handler(dto.SignUpRequest{
		Email:    "qwerty",
		Password: "123",
	})

	assert.NoError(t, actualErr)
	assert.Equal(t, expectedResponse, actualResponse)
}

func TestSignUpHandler_Negative(t *testing.T) {
	service := &mockService{fnR: func(req dto.SignUpRequest) (*model.User, error) {
		assert.Equal(t, "qwerty", req.Email)
		assert.Equal(t, "123", req.Password)

		return nil, errors.New("fuck up")
	}}

	handler := SignUpHandler(service)

	_, actualErr := handler(dto.SignUpRequest{
		Email:    "qwerty",
		Password: "123",
	})

	assert.Error(t, actualErr)
}
