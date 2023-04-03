package handlers

import (
	"bytes"
	"context"
	"drunklish/internal/config"
	"drunklish/internal/connection"
	"drunklish/internal/model"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth"
	"drunklish/internal/service/auth/dto"
	"drunklish/internal/service/auth/repository"
	"drunklish/internal/service/auth/token"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignInHandler_Positive(t *testing.T) {
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

	reqIn := httptest.NewRequest("POST", "/sign-in", bytes.NewBuffer([]byte(`{"email":"asd@gmail.com","hash_password":"qwerty"}`)))
	resIn := httptest.NewRecorder()

	handlerIn := SignInHandler(authDB)
	handlerIn(resIn, reqIn)
	authUser, err := authDB.SignIn(model.User{
		Email:        "asd@gmail.com",
		HashPassword: "qwerty",
	})
	assert.NoError(t, err)

	hashPassword, err := token.HashPassword(authUser.User.HashPassword)
	assert.NoError(t, err)
	authUser.User.HashPassword = hashPassword

	assert.Equal(t, &dto.ResponseUser{
		User: model.User{
			Id:           authUser.User.Id,
			Email:        "asd@gmail.com",
			HashPassword: authUser.User.HashPassword,
		},
		Token: authUser.Token,
	}, authUser)
}

func TestSignInHandler_NegativeFailSignIn(t *testing.T) {
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

	reqIn := httptest.NewRequest("POST", "/sign-in", bytes.NewBuffer([]byte(`{"email":"asd@gmail.com.com","hash_password":"qwerty"}`)))
	resIn := httptest.NewRecorder()

	handlerIn := SignInHandler(authDB)
	handlerIn(resIn, reqIn)

	_, err = authDB.SignIn(model.User{
		Email:        "asd@gmail.com.com",
		HashPassword: "qwerty",
	})
	assert.ErrorIs(t, err, httputils.ErrValidation)
	assert.Equal(t, http.StatusBadRequest, resIn.Code)
}

func TestSignInHandler_NegativeFailUnmarshal(t *testing.T) {
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

	reqIn := httptest.NewRequest("POST", "/sign-in", bytes.NewBuffer([]byte(`{"email":"asd@gmail.com","hash_password":"qwerty",}`)))
	resIn := httptest.NewRecorder()

	handlerIn := SignInHandler(authDB)
	handlerIn(resIn, reqIn)

	//assert.ErrorIs(t, err, httputils.UnmarshalError)
	assert.Equal(t, http.StatusBadRequest, resIn.Code)
}
