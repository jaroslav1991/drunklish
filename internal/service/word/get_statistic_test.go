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

func TestWord_GetStatistic_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trainingInfo := dto.RequestStatistic{
		TrainingId: 1,
		Token:      "token",
	}

	responseTraining := dto.ResponseTrainingInfo{
		Words: dto.ResponseWordList{
			{
				Id:        1,
				Word:      "test",
				Translate: "test",
			},
			{
				Id:        2,
				Word:      "test2",
				Translate: "test2",
			},
		},
		Answers: dto.ResponseWordList{
			{
				Id:        1,
				Word:      "test",
				Translate: "test",
			},
			{
				Id:        2,
				Word:      "test2",
				Translate: "test",
			},
		},
	}

	statistic := dto.RequestStatistic{
		TrainingId: 1,
		Token:      "token",
	}

	responseStatistic := dto.ResponseStatistic{
		CorrectAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        1,
			Word:      "test",
			Translate: "test",
		}}},
		WrongAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        2,
			Word:      "test2",
			Translate: "test",
		}}},
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().GetTrainingInfoById(trainingInfo.TrainingId, int64(1)).Return(&responseTraining, nil)
	repository.EXPECT().GetStatisticByTrainingId(statistic.TrainingId, int64(1)).Return(&responseStatistic, nil)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	actualStatistic, err := service.GetStatistic(statistic)
	assert.NoError(t, err)

	assert.Equal(t, &dto.ResponseStatistic{
		CorrectAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        1,
			Word:      "test",
			Translate: "test",
		}}},
		WrongAnswers: dto.ResponseWords{Words: []dto.ResponseWord{{
			Id:        2,
			Word:      "test2",
			Translate: "test",
		}}},
	}, actualStatistic)
}

func TestWord_GetStatistic_NegativeFailStatistic(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trainingInfo := dto.RequestStatistic{
		TrainingId: 1,
		Token:      "token",
	}

	responseTraining := dto.ResponseTrainingInfo{
		Words: dto.ResponseWordList{
			{
				Id:        1,
				Word:      "test",
				Translate: "test",
			},
			{
				Id:        2,
				Word:      "test2",
				Translate: "test2",
			},
		},
		Answers: dto.ResponseWordList{
			{
				Id:        1,
				Word:      "test",
				Translate: "test",
			},
			{
				Id:        2,
				Word:      "test2",
				Translate: "test",
			},
		},
	}

	statistic := dto.RequestStatistic{
		TrainingId: 1,
		Token:      "token",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().GetTrainingInfoById(trainingInfo.TrainingId, int64(1)).Return(&responseTraining, nil)
	repository.EXPECT().GetStatisticByTrainingId(statistic.TrainingId, int64(1)).Return(nil, errors.New("fail get statistic"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.GetStatistic(statistic)
	assert.Error(t, err)
}

func TestWord_GetStatistic_NegativeFailTraining(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trainingInfo := dto.RequestStatistic{
		TrainingId: 1,
		Token:      "token",
	}

	statistic := dto.RequestStatistic{
		TrainingId: 1,
		Token:      "token",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().GetTrainingInfoById(trainingInfo.TrainingId, int64(1)).Return(nil, errors.New("fail get training"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.GetStatistic(statistic)
	assert.Error(t, err)
}

func TestWord_GetStatistic_NegativeFailToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statistic := dto.RequestStatistic{
		TrainingId: 1,
		Token:      "token",
	}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return nil, errors.New("fuck up")
	}

	_, err := service.GetStatistic(statistic)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}
