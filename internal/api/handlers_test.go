package api

import (
	"bytes"
	"context"
	"drunklish/internal/config"
	"drunklish/internal/model"
	"drunklish/internal/service/auth"
	"drunklish/internal/service/auth/users"
	"drunklish/pkg/repository"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestSignUpHandlerPositive(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		t.Error(err)
	}

	tx, err := db.BeginTxx(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	defer tx.Rollback()

	if _, err := tx.Exec("drop table words"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("drop table users"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table words (id bigserial primary key,word varchar(55) not null,translate varchar(55) not null,created_at timestamp,user_id bigint references users(id))"); err != nil {
		t.Error(err)
	}

	authDB := auth.NewAuthService(tx)

	req := httptest.NewRequest("POST", "/sign-up", bytes.NewBuffer([]byte(`{"email":"test@gmail.com","hash_password":"qwerty"}`)))
	res := httptest.NewRecorder()

	handler := SignUpHandler(authDB)
	handler(res, req)
}

func TestSignInHandlerNegative(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		t.Error(err)
	}

	tx, err := db.BeginTxx(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	defer tx.Rollback()

	if _, err := tx.Exec("drop table words"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("drop table users"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table words (id bigserial primary key,word varchar(55) not null,translate varchar(55) not null,created_at timestamp,user_id bigint references users(id))"); err != nil {
		t.Error(err)
	}

	authDB := auth.NewAuthService(tx)

	userDB, err := authDB.SignUp(&model.User{
		Email:        "test@gmail.com",
		HashPassword: "qwerty",
	})

	if err != nil {
		t.Error(err)
	}
	hashPassword, err := users.HashPassword(userDB.HashPassword)
	if err != nil {
		t.Error(err)
	}
	userDB.HashPassword = hashPassword

	marshal, err := json.Marshal(userDB)
	if err != nil {
		t.Error()
	}

	req := httptest.NewRequest("POST", "/sign-up", bytes.NewBuffer(marshal))
	res := httptest.NewRecorder()

	handler := SignUpHandler(authDB)
	handler(res, req)
}

func TestSignInHandlerUnmarshalNegative(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		t.Error(err)
	}

	tx, err := db.BeginTxx(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	defer tx.Rollback()

	if _, err := tx.Exec("drop table words"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("drop table users"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table words (id bigserial primary key,word varchar(55) not null,translate varchar(55) not null,created_at timestamp,user_id bigint references users(id))"); err != nil {
		t.Error(err)
	}

	authDB := auth.NewAuthService(tx)

	req := httptest.NewRequest("POST", "/sign-up", nil)
	res := httptest.NewRecorder()

	handler := SignUpHandler(authDB)
	handler(res, req)
}

func TestSignInHandlerErrorDomain(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		t.Error(err)
	}

	tx, err := db.BeginTxx(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	defer tx.Rollback()

	if _, err := tx.Exec("drop table words"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("drop table users"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table words (id bigserial primary key,word varchar(55) not null,translate varchar(55) not null,created_at timestamp,user_id bigint references users(id))"); err != nil {
		t.Error(err)
	}

	authDB := auth.NewAuthService(tx)

	_, err = authDB.SignUp(&model.User{
		Email:        "test@gmail.ru",
		HashPassword: "qwerty",
	})

	assert.ErrorIs(t, err, auth.ErrDomain)
}

func TestSignInHandlerErrorSymbol(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		t.Error(err)
	}

	tx, err := db.BeginTxx(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	defer tx.Rollback()

	if _, err := tx.Exec("drop table words"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("drop table users"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table words (id bigserial primary key,word varchar(55) not null,translate varchar(55) not null,created_at timestamp,user_id bigint references users(id))"); err != nil {
		t.Error(err)
	}

	authDB := auth.NewAuthService(tx)

	_, err = authDB.SignUp(&model.User{
		Email:        "test@@gmail.com",
		HashPassword: "qwerty",
	})

	assert.ErrorIs(t, err, auth.ErrSymbol)
}

func TestSignInHandlerErrorLengthPassword(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		t.Error(err)
	}

	tx, err := db.BeginTxx(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	defer tx.Rollback()

	if _, err := tx.Exec("drop table words"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("drop table users"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table words (id bigserial primary key,word varchar(55) not null,translate varchar(55) not null,created_at timestamp,user_id bigint references users(id))"); err != nil {
		t.Error(err)
	}

	authDB := auth.NewAuthService(tx)

	_, err = authDB.SignUp(&model.User{
		Email:        "test@gmail.com",
		HashPassword: "1234",
	})

	assert.ErrorIs(t, err, auth.ErrLengthPassword)
}

func TestSignInHandlerErrorExistEmail(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		t.Error(err)
	}

	tx, err := db.BeginTxx(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	defer tx.Rollback()

	if _, err := tx.Exec("drop table words"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("drop table users"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)"); err != nil {
		t.Error(err)
	}

	if _, err := tx.Exec("create table words (id bigserial primary key,word varchar(55) not null,translate varchar(55) not null,created_at timestamp,user_id bigint references users(id))"); err != nil {
		t.Error(err)
	}

	authDB := auth.NewAuthService(tx)

	_, err = authDB.SignUp(&model.User{
		Email:        "test@gmail.com",
		HashPassword: "qwerty",
	})
	assert.NoError(t, err)

	_, err = authDB.SignUp(&model.User{
		Email:        "test@gmail.com",
		HashPassword: "qwerty",
	})
	assert.ErrorIs(t, err, auth.ErrExistEmail)
}
