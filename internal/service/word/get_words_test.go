package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth/token"
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestWord_GetWordsByUserId_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tkn := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2LCJFbWFpbCI6InF3ZUBnbWFpbC5jb20iLCJleHAiOjE2ODE0MTQxMzR9.gq-wuyDVK9enmmp9yWXeKM9JM2YYKxJKcsY1LeQuflM"

	word := dto.RequestForGettingWord{Token: tkn}

	wordsFromGet := dto.ResponseWords{Words: []dto.ResponseWord{{
		Word:      "boogaga1",
		Translate: "boo1",
	}, {
		Word:      "boogaga2",
		Translate: "boo",
	}}}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().CheckUserInDB(int64(1)).Return(true, nil)
	repository.EXPECT().GetWords(int64(1)).Return(&wordsFromGet, nil)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	actualWords, err := service.GetWordsByUserId(word)
	assert.NoError(t, err)

	assert.Equal(t, &dto.ResponseWords{Words: []dto.ResponseWord{{
		Word:      "boogaga1",
		Translate: "boo1",
	}, {
		Word:      "boogaga2",
		Translate: "boo",
	}}}, actualWords)
}

func TestWord_GetWordsByUserId_NegativeFailGetWord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tkn := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2LCJFbWFpbCI6InF3ZUBnbWFpbC5jb20iLCJleHAiOjE2ODE0MTQxMzR9.gq-wuyDVK9enmmp9yWXeKM9JM2YYKxJKcsY1LeQuflM"

	word := dto.RequestForGettingWord{Token: tkn}
	repository := NewMockRepository(ctrl)

	repository.EXPECT().CheckUserInDB(int64(1)).Return(true, nil)
	repository.EXPECT().GetWords(int64(1)).Return(nil, errors.New("fail create word"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.GetWordsByUserId(word)
	assert.ErrorIs(t, err, httputils.ErrInternalServer)

}

func TestWord_GetWordsByUserId_NegativeFailCheckUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tkn := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2LCJFbWFpbCI6InF3ZUBnbWFpbC5jb20iLCJleHAiOjE2ODE0MTQxMzR9.gq-wuyDVK9enmmp9yWXeKM9JM2YYKxJKcsY1LeQuflM"

	word := dto.RequestForGettingWord{Token: tkn}
	repository := NewMockRepository(ctrl)

	repository.EXPECT().CheckUserInDB(int64(1)).Return(true, errors.New("fuck up"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.GetWordsByUserId(word)
	assert.ErrorIs(t, err, httputils.ErrValidation)

}

func TestWord_GetWordsByUserId_NegativeFailNotCheckUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tkn := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2LCJFbWFpbCI6InF3ZUBnbWFpbC5jb20iLCJleHAiOjE2ODE0MTQxMzR9.gq-wuyDVK9enmmp9yWXeKM9JM2YYKxJKcsY1LeQuflM"

	word := dto.RequestForGettingWord{Token: tkn}
	repository := NewMockRepository(ctrl)

	repository.EXPECT().CheckUserInDB(int64(1)).Return(false, errors.New("fuck up"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.GetWordsByUserId(word)
	assert.ErrorIs(t, err, httputils.ErrValidation)

}

func TestWord_GetWordsByUserId_NegativeFailToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tkn := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2LCJFbWFpbCI6InF3ZUBnbWFpbC5jb20iLCJleHAiOjE2ODE0MTQxMzR9.gq-wuyDVK9enmmp9yWXeKM9JM2YYKxJKcsY1LeQuflM"

	word := dto.RequestForGettingWord{Token: tkn}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return nil, errors.New("fuck up")
	}

	_, err := service.GetWordsByUserId(word)
	assert.Error(t, err)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}

func TestWord_GetWordsByCreatedAt_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := dto.RequestForGetByPeriod{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2LCJFbWFpbCI6InF3ZUBnbWFpbC5jb20iLCJleHAiOjE2ODE0MTQxMzR9.gq-wuyDVK9enmmp9yWXeKM9JM2YYKxJKcsY1LeQuflM",
		CreatedAt: dto.Period{
			FirstDate:  time.Now(),
			SecondDate: time.Now(),
		},
	}

	wordsFromGet := dto.ResponseWords{Words: []dto.ResponseWord{{
		Word:      "boogaga1",
		Translate: "boo1",
	}, {
		Word:      "boogaga2",
		Translate: "boo",
	}}}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().CheckUserInDB(int64(1)).Return(true, nil)
	repository.EXPECT().GetWordByCreated(int64(1), word.CreatedAt.FirstDate, word.CreatedAt.SecondDate).Return(&wordsFromGet, nil)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	actualWords, err := service.GetWordsByCreatedAt(word)
	assert.NoError(t, err)

	assert.Equal(t, &dto.ResponseWords{Words: []dto.ResponseWord{{
		Word:      "boogaga1",
		Translate: "boo1",
	}, {
		Word:      "boogaga2",
		Translate: "boo",
	}}}, actualWords)
}

func TestWord_GetWordsByCreatedAt_NegativeFailGetWord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := dto.RequestForGetByPeriod{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2LCJFbWFpbCI6InF3ZUBnbWFpbC5jb20iLCJleHAiOjE2ODE0MTQxMzR9.gq-wuyDVK9enmmp9yWXeKM9JM2YYKxJKcsY1LeQuflM", CreatedAt: dto.Period{
			FirstDate:  time.Now(),
			SecondDate: time.Now(),
		},
	}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().CheckUserInDB(int64(1)).Return(true, nil)
	repository.EXPECT().GetWordByCreated(int64(1), word.CreatedAt.FirstDate, word.CreatedAt.SecondDate).Return(nil, errors.New("fuck up"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.GetWordsByCreatedAt(word)
	assert.Error(t, err)
	assert.ErrorIs(t, err, httputils.ErrInternalServer)
}

func TestWord_GetWordsByCreatedAt_NegativeFailCheckUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := dto.RequestForGetByPeriod{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2LCJFbWFpbCI6InF3ZUBnbWFpbC5jb20iLCJleHAiOjE2ODE0MTQxMzR9.gq-wuyDVK9enmmp9yWXeKM9JM2YYKxJKcsY1LeQuflM", CreatedAt: dto.Period{
			FirstDate:  time.Now(),
			SecondDate: time.Now(),
		},
	}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().CheckUserInDB(int64(1)).Return(true, errors.New("fuck up"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.GetWordsByCreatedAt(word)
	assert.Error(t, err)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}

func TestWord_GetWordsByCreatedAt_NegativeFailNotCheckUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := dto.RequestForGetByPeriod{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2LCJFbWFpbCI6InF3ZUBnbWFpbC5jb20iLCJleHAiOjE2ODE0MTQxMzR9.gq-wuyDVK9enmmp9yWXeKM9JM2YYKxJKcsY1LeQuflM", CreatedAt: dto.Period{
			FirstDate:  time.Now(),
			SecondDate: time.Now(),
		},
	}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().CheckUserInDB(int64(1)).Return(false, errors.New("fuck up"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.GetWordsByCreatedAt(word)
	assert.Error(t, err)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}

func TestWord_GetWordsByCreatedAt_NegativeFailToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := dto.RequestForGetByPeriod{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjE2LCJFbWFpbCI6InF3ZUBnbWFpbC5jb20iLCJleHAiOjE2ODE0MTQxMzR9.gq-wuyDVK9enmmp9yWXeKM9JM2YYKxJKcsY1LeQuflM", CreatedAt: dto.Period{
			FirstDate:  time.Now(),
			SecondDate: time.Now(),
		},
	}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return nil, errors.New("fuck up")
	}

	_, err := service.GetWordsByCreatedAt(word)
	assert.Error(t, err)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}
