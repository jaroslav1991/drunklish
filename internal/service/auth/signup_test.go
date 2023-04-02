package auth

import (
	"drunklish/internal/model"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth/dto"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuth_SignUp_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userFromRepository := model.User{
		Id:           1,
		Email:        "new@gmail.com",
		HashPassword: "qwerty",
	}

	dtoForCreateUser := dto.SignUpRequest{
		Email:    "new@gmail.com",
		Password: "hash",
	}

	requestDto := dto.SignUpRequest{
		Email:    "new@gmail.com",
		Password: "qwerty",
	}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().ExistEmail("new@gmail.com").Return(false, nil)
	repository.EXPECT().CreateUser(dtoForCreateUser).Return(&userFromRepository, nil)

	service := NewAuthService(repository)
	service.hashFn = func(password string) (string, error) {
		return "hash", nil
	}

	actualUser, err := service.SignUp(requestDto)
	assert.NoError(t, err)

	assert.Equal(t, model.User{
		Id:           1,
		Email:        "new@gmail.com",
		HashPassword: "qwerty",
	}, *actualUser)
}

func TestAuth_SignUp_Negative_FailToCreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	dtoForCreateUser := dto.SignUpRequest{
		Email:    "new@gmail.com",
		Password: "hash",
	}

	requestDto := dto.SignUpRequest{
		Email:    "new@gmail.com",
		Password: "qwerty",
	}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().ExistEmail("new@gmail.com").Return(false, nil)
	repository.EXPECT().CreateUser(dtoForCreateUser).Return(nil, errors.New("fail on create user"))

	service := NewAuthService(repository)
	service.hashFn = func(password string) (string, error) {
		return "hash", nil
	}

	_, err := service.SignUp(requestDto)
	assert.Error(t, err)
}

func TestAuth_SignUp_Negative_FailHashPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestDto := dto.SignUpRequest{
		Email:    "new@gmail.com",
		Password: "qwerty",
	}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().ExistEmail("new@gmail.com").Return(false, nil)

	service := NewAuthService(repository)
	service.hashFn = func(password string) (string, error) {
		return "", errors.New("fuck up")
	}

	_, err := service.SignUp(requestDto)
	assert.ErrorIs(t, err, httputils.ErrInternalServer)
}

func TestAuth_SignUp_Negative_FailExistEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestDto := dto.SignUpRequest{
		Email:    "new@gmail.com",
		Password: "qwerty",
	}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().ExistEmail("new@gmail.com").Return(true, errors.New("fail email exist"))

	service := NewAuthService(repository)

	_, err := service.SignUp(requestDto)
	assert.ErrorIs(t, err, httputils.ErrExistEmail)
}

func TestAuth_SignUp_Negative_FailEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestDto := dto.SignUpRequest{
		Email:    "new@gmail.com",
		Password: "qwerty",
	}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().ExistEmail("new@gmail.com").Return(false, errors.New("check exist email"))

	service := NewAuthService(repository)

	_, err := service.SignUp(requestDto)
	assert.ErrorIs(t, err, httputils.ErrInternalServer)
}

func TestAuth_SignUp_Negative_FailLengthPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestDto := dto.SignUpRequest{
		Email:    "new@gmail.com",
		Password: "123",
	}

	repository := NewMockRepository(ctrl)

	service := NewAuthService(repository)

	_, err := service.SignUp(requestDto)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}

func TestAuth_SignUp_Negative_FailSymbol(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestDto := dto.SignUpRequest{
		Email:    "new@@gmail.com",
		Password: "qwerty",
	}

	repository := NewMockRepository(ctrl)

	service := NewAuthService(repository)

	_, err := service.SignUp(requestDto)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}

func TestAuth_SignUp_Negative_FailDomain(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	requestDto := dto.SignUpRequest{
		Email:    "new@google.ru",
		Password: "qwerty",
	}

	repository := NewMockRepository(ctrl)

	service := NewAuthService(repository)

	_, err := service.SignUp(requestDto)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}
