package word

import (
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth/token"
	"drunklish/internal/service/word/dto"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWord_CreateStatistic_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createdStatistic := dto.RequestCreateStatistic{
		TrainingId: 1,
		CorrectAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        1,
			Word:      "test",
			Translate: "test",
		}, {
			Id:        2,
			Word:      "test2",
			Translate: "test2",
		},
		}},
		WrongAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        3,
			Word:      "test3",
			Translate: "qwe3",
		}}},
		Token: "token",
	}

	responseStatistic := dto.ResponseCreateStatistic{Id: int64(1)}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().CreateStatisticByTrainingId(
		createdStatistic.TrainingId,
		createdStatistic.CorrectAnswers.Words,
		createdStatistic.WrongAnswers.Words,
		int64(1),
	).Return(&responseStatistic, nil)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	actualStatistic, err := service.CreateStatistic(createdStatistic)
	assert.NoError(t, err)

	assert.Equal(t, &dto.ResponseCreateStatistic{Id: int64(1)}, actualStatistic)
}

func TestWord_CreateStatistic_NegativeFailCreateStatistic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createdStatistic := dto.RequestCreateStatistic{
		TrainingId: 1,
		CorrectAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        1,
			Word:      "test",
			Translate: "test",
		}, {
			Id:        2,
			Word:      "test2",
			Translate: "test2",
		},
		}},
		WrongAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        3,
			Word:      "test3",
			Translate: "qwe3",
		}}},
		Token: "token",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().CreateStatisticByTrainingId(
		createdStatistic.TrainingId,
		createdStatistic.CorrectAnswers.Words,
		createdStatistic.WrongAnswers.Words,
		int64(1),
	).Return(nil, errors.New("fail create statistic"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.CreateStatistic(createdStatistic)
	assert.Error(t, err)
}

func TestWord_CreateStatistic_NegativeFailToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	createdStatistic := dto.RequestCreateStatistic{
		TrainingId: 1,
		CorrectAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        1,
			Word:      "test",
			Translate: "test",
		}, {
			Id:        2,
			Word:      "test2",
			Translate: "test2",
		},
		}},
		WrongAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        3,
			Word:      "test3",
			Translate: "qwe3",
		}}},
		Token: "token",
	}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return nil, errors.New("fuck up")
	}

	_, err := service.CreateStatistic(createdStatistic)
	assert.ErrorIs(t, err, httputils.ErrValidation)

}
