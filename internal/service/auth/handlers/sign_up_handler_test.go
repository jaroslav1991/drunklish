package handlers

import (
	"bytes"
	"context"
	"drunklish/internal/config"
	"drunklish/internal/connection"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth"
	"drunklish/internal/service/auth/dto"
	"drunklish/internal/service/auth/repository"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignUpHandler_Positive(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := connection.NewPostgresDB(dbConfig)
	assert.NoError(t, err)

	tx, err := db.BeginTxx(context.Background(), nil)
	assert.NoError(t, err)

	defer tx.Rollback()

	_, err = tx.Exec("create table if not exists users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)")
	assert.NoError(t, err)

	repo := repository.NewAuthRepository(tx)
	authDB := auth.NewAuthService(repo)

	req := httptest.NewRequest("POST", "/sign-up", bytes.NewBuffer([]byte(`{"email":"asd@gmail.com","password":"qwerty"}`)))
	res := httptest.NewRecorder()

	handler := SignUpHandler(authDB)
	handler(res, req)

	assert.Equal(t, http.StatusOK, res.Code)
}

func TestSignUpHandler_NegativeFailSignUp(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := connection.NewPostgresDB(dbConfig)
	assert.NoError(t, err)

	tx, err := db.BeginTxx(context.Background(), nil)
	assert.NoError(t, err)

	defer tx.Rollback()

	_, err = tx.Exec("create table if not exists users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)")
	assert.NoError(t, err)

	repo := repository.NewAuthRepository(tx)
	authDB := auth.NewAuthService(repo)

	req := httptest.NewRequest("POST", "/sign-up", bytes.NewBuffer([]byte(`{"email":"asd@gyandex.com.com","password":"qwerty"}`)))
	res := httptest.NewRecorder()

	handler := SignUpHandler(authDB)
	handler(res, req)

	_, err = authDB.SignUp(dto.SignUpRequest{
		Email:    "asd@gyandex.com",
		Password: "qwerty",
	})

	assert.ErrorIs(t, err, httputils.ErrValidation)
	assert.Equal(t, http.StatusBadRequest, res.Code)
}

func TestSignUpHandler_NegativeFailUnmarshal(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := connection.NewPostgresDB(dbConfig)
	assert.NoError(t, err)

	tx, err := db.BeginTxx(context.Background(), nil)
	assert.NoError(t, err)

	defer tx.Rollback()

	_, err = tx.Exec("create table if not exists users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)")
	assert.NoError(t, err)

	repo := repository.NewAuthRepository(tx)
	authDB := auth.NewAuthService(repo)

	req := httptest.NewRequest("POST", "/sign-up", bytes.NewBuffer([]byte(`{"email":"asd@yandex.ru","password":"qwerty",}`)))
	res := httptest.NewRecorder()

	handler := SignUpHandler(authDB)
	handler(res, req)

	assert.Equal(t, http.StatusBadRequest, res.Code)
}
