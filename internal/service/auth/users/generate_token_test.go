package users

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password1, err := HashPassword("1234")
	assert.NoError(t, err)

	password2, err := HashPassword("1234")
	assert.NoError(t, err)

	t.Log(password1)
	t.Log(password2)

}

func TestCheckPasswordHash(t *testing.T) {
	hash, err := HashPassword("1234")
	assert.NoError(t, err)

	t.Log(hash)
	result := CheckPasswordHash("1234", hash)
	assert.NoError(t, result)
}

func TestGenerateToken(t *testing.T) {
	generate, err := GenerateToken(1, "test@mail.ru")
	assert.NoError(t, err)
	t.Log(generate)

	token, err := ParseToken(generate)
	assert.NoError(t, err)
	t.Log(token)

	_, err = GenerateToken(0, "test@mail.ru")
	assert.ErrorIs(t, err, InvalidToken)
}

func TestParseToken(t *testing.T) {
	gen, err := GenerateToken(2, "vasya")
	assert.NoError(t, err)

	ParseToken(gen)
}
