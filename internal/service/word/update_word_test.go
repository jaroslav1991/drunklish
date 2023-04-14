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

func TestWord_UpdateWord_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := dto.RequestForUpdateWord{
		Word:      "qwe1",
		Translate: "qwe1",
		Id:        1,
		Token:     "qwerty",
	}
	wordFromUpdate := dto.ResponseWord{
		Id:        1,
		Word:      "qwe2",
		Translate: "qwe2",
	}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().Update(word.Word, word.Translate, word.Id, int64(1)).Return(&wordFromUpdate, nil)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	actual, err := service.UpdateWord(word)
	assert.NoError(t, err)
	assert.Equal(t, &dto.ResponseWord{
		Id:        1,
		Word:      "qwe2",
		Translate: "qwe2",
	}, actual)
}

func TestWord_UpdateWord_NegativeFailUpdateWord(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := dto.RequestForUpdateWord{
		Word:      "qwe1",
		Translate: "qwe1",
		Id:        1,
		Token:     "qwerty",
	}

	repository := NewMockRepository(ctrl)

	repository.EXPECT().Update(word.Word, word.Translate, word.Id, int64(1)).Return(nil, errors.New("fuck up"))

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return &token.AuthClaims{UserId: int64(1)}, nil
	}

	_, err := service.UpdateWord(word)
	assert.Error(t, err)
}

func TestWord_UpdateWord_NegativeFailToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	word := dto.RequestForUpdateWord{
		Word:      "qwe1",
		Translate: "qwe1",
		Id:        1,
		Token:     "qwerty",
	}

	repository := NewMockRepository(ctrl)

	service := NewWordService(repository)
	service.parseTokenFn = func(tokenString string) (*token.AuthClaims, error) {
		return nil, errors.New("fuck up")
	}

	_, err := service.UpdateWord(word)
	assert.Error(t, err)
	assert.ErrorIs(t, err, httputils.ErrValidation)
}
