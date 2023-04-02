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

func TestAuth_SignIn_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userModel := model.User{
		Email:        "new@gmail.com",
		HashPassword: "qwerty",
	}

	userFromDB := dto.ResponseUser{
		User:  userModel,
		Token: "token",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().ExistEmail(userModel.Email).Return(true, nil)
	repository.EXPECT().CheckUserDB(userModel).Return(&userFromDB, nil)

	service := NewAuthService(repository)
	service.checkPasswordFn = func(hash, password string) error {
		return nil
	}
	service.generateTokenFn = func(userId int64, email string) (string, error) {
		return "token", nil
	}

	actual, err := service.SignIn(userModel)
	assert.NoError(t, err)

	assert.Equal(t, userFromDB, *actual)
}

func TestAuth_SignIn_Negative_FailGenerateToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userModel := model.User{
		Email:        "new@gmail.com",
		HashPassword: "qwerty",
	}

	userFromDB := dto.ResponseUser{
		User:  userModel,
		Token: "token",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().ExistEmail(userModel.Email).Return(true, nil)
	repository.EXPECT().CheckUserDB(userModel).Return(&userFromDB, nil)

	service := NewAuthService(repository)
	service.checkPasswordFn = func(hash, password string) error {
		return nil
	}
	service.generateTokenFn = func(userId int64, email string) (string, error) {
		return "", errors.New("fail generate token")
	}

	_, err := service.SignIn(userModel)
	assert.ErrorIs(t, err, httputils.ErrInternalServer)
}

func TestAuth_SignIn_Negative_FailCheckPassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userModel := model.User{
		Email:        "new@gmail.com",
		HashPassword: "qwerty",
	}

	userFromDB := dto.ResponseUser{
		User:  userModel,
		Token: "token",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().ExistEmail(userModel.Email).Return(true, nil)
	repository.EXPECT().CheckUserDB(userModel).Return(&userFromDB, nil)

	service := NewAuthService(repository)
	service.checkPasswordFn = func(hash, password string) error {
		return errors.New("fail check password")
	}

	_, err := service.SignIn(userModel)
	assert.ErrorIs(t, err, httputils.ErrInternalServer)
}

func TestAuth_SignIn_Negative_FailCheckUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userModel := model.User{
		Email:        "new@gmail.com",
		HashPassword: "qwerty",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().ExistEmail(userModel.Email).Return(true, nil)
	repository.EXPECT().CheckUserDB(userModel).Return(nil, errors.New("fail check user DB"))

	service := NewAuthService(repository)

	_, err := service.SignIn(userModel)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}

func TestAuth_SignIn_Negative_FailExistEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userModel := model.User{
		Email:        "new@gmail.com",
		HashPassword: "qwerty",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().ExistEmail(userModel.Email).Return(true, errors.New(""))

	service := NewAuthService(repository)

	_, err := service.SignIn(userModel)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}

func TestAuth_SignIn_Negative_FailNotExistEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userModel := model.User{
		Email:        "new@gmail.com",
		HashPassword: "qwerty",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().ExistEmail(userModel.Email).Return(false, errors.New(""))

	service := NewAuthService(repository)

	_, err := service.SignIn(userModel)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}
