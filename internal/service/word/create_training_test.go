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

func TestWord_CreateTrainingWords_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trainingForCreate := dto.RequestForTraining{
		Words:      dto.ResponseWords{Words: []dto.ResponseWord{{Id: 1, Word: "test", Translate: "test"}}},
		Answers:    dto.ResponseWords{Words: []dto.ResponseWord{{Id: 1, Word: "test", Translate: "test"}}},
		WordsTotal: 1,
		Token:      "token",
	}
	trainingResponse := dto.ResponseForTraining{Id: int64(1)}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().CreateTraining(
		trainingForCreate.Words.Words,
		trainingForCreate.Answers.Words,
		trainingForCreate.WordsTotal,
		int64(1)).Return(&trainingResponse, nil)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	actualTraining, err := service.CreateTrainingWords(trainingForCreate)
	assert.NoError(t, err)

	assert.Equal(t, &dto.ResponseForTraining{Id: int64(1)}, actualTraining)
}

func TestWord_CreateTrainingWords_NegativeFailCreateTraining(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trainingForCreate := dto.RequestForTraining{
		Words:      dto.ResponseWords{Words: []dto.ResponseWord{{Id: 1, Word: "test", Translate: "test"}}},
		Answers:    dto.ResponseWords{Words: []dto.ResponseWord{{Id: 1, Word: "test", Translate: "test"}}},
		WordsTotal: 1,
		Token:      "token",
	}

	repository := NewMockRepository(ctrl)
	repository.EXPECT().CreateTraining(
		trainingForCreate.Words.Words,
		trainingForCreate.Answers.Words,
		trainingForCreate.WordsTotal,
		int64(1)).Return(nil, errors.New("fail create training"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.CreateTrainingWords(trainingForCreate)
	assert.Error(t, err)
}

func TestWord_CreateTrainingWords_NegativeFailToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	trainingForCreate := dto.RequestForTraining{
		Words:      dto.ResponseWords{Words: []dto.ResponseWord{{Id: 1, Word: "test", Translate: "test"}}},
		Answers:    dto.ResponseWords{Words: []dto.ResponseWord{{Id: 1, Word: "test", Translate: "test"}}},
		WordsTotal: 1,
		Token:      "token",
	}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return nil, errors.New("fuck up")
	}

	_, err := service.CreateTrainingWords(trainingForCreate)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}
