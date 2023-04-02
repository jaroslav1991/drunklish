package auth

import (
	"drunklish/internal/model"
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
