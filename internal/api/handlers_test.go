package api

import (
	"bytes"
	"context"
	"drunklish/internal/config"
	"drunklish/internal/model"
	"drunklish/internal/service/auth"
	"drunklish/internal/service/auth/users"
	"drunklish/internal/service/word"
	"drunklish/pkg/repository"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

// Testing SignUp ---------------------------------------------------------------------------------------------------

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

func TestSignUpHandlerNegative(t *testing.T) {
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

func TestSignUpHandlerUnmarshalNegative(t *testing.T) {
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

func TestSignUpHandlerErrorDomain(t *testing.T) {
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

func TestSignUpHandlerErrorSymbol(t *testing.T) {
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

func TestSignUpHandlerErrorLengthPassword(t *testing.T) {
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

func TestSignUpHandlerErrorExistEmail(t *testing.T) {
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

//Testing SignIn --------------------------------------------------------------------------------------------------

func TestSignInHandlerPositive(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	assert.NoError(t, err)

	authDB := auth.NewAuthService(db)

	req := httptest.NewRequest("POST", "/sign-in", bytes.NewBuffer([]byte(`{"email":"test@yahoo.com","hash_password":"qwerty"}`)))
	res := httptest.NewRecorder()

	handler := SignInHandler(authDB)
	handler(res, req)

	userDB, err := authDB.SignIn(&model.User{
		Email:        "test@yahoo.com",
		HashPassword: "qwerty",
	})
	assert.NoError(t, err)

	hashPassword, err := users.HashPassword(userDB.User.HashPassword)
	assert.NoError(t, err)
	userDB.User.HashPassword = hashPassword

	assert.Equal(t, &auth.ResponseUser{
		User: model.User{
			Id:           2,
			Email:        "test@yahoo.com",
			HashPassword: userDB.User.HashPassword,
		},
		Token: userDB.Token,
	}, userDB)
}

func TestSignInHandlerNegativeErrEmail(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	assert.NoError(t, err)

	authDB := auth.NewAuthService(db)

	req := httptest.NewRequest("POST", "/sign-in", bytes.NewBuffer([]byte(`{"email":"lox@yahoo.com","hash_password":"qwerty"}`)))
	res := httptest.NewRecorder()

	handler := SignInHandler(authDB)
	handler(res, req)

	_, err = authDB.SignIn(&model.User{
		Email:        "lox@yahoo.com",
		HashPassword: "qwerty",
	})
	assert.ErrorIs(t, err, auth.ErrEmail)
}

func TestSignInHandlerNegativeErrCheckPassword(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	assert.NoError(t, err)

	authDB := auth.NewAuthService(db)

	req := httptest.NewRequest("POST", "/sign-in", bytes.NewBuffer([]byte(`{"email":"test@yahoo.com","hash_password":"qwerty123"}`)))
	res := httptest.NewRecorder()

	handler := SignInHandler(authDB)
	handler(res, req)

	_, err = authDB.SignIn(&model.User{
		Email:        "test@yahoo.com",
		HashPassword: "qwerty123",
	})

	assert.ErrorIs(t, err, auth.ErrPassword)
}

func TestSignInHandlerNegativeErrUnmarshal(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	assert.NoError(t, err)

	authDB := auth.NewAuthService(db)

	req := httptest.NewRequest("POST", "/sign-in", nil)
	res := httptest.NewRecorder()

	handler := SignInHandler(authDB)
	handler(res, req)
}

func TestSignInHandlerNegativeErrToken(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	assert.NoError(t, err)

	auth.NewAuthService(db)

	var user model.User

	_, err = users.GenerateToken(user.Id, user.Email)
	assert.ErrorIs(t, err, users.InvalidToken)
}

// Testing DeleteWord ----------------------------------------------------------------------------------------------

func TestDeleteWordHandlerPositive(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	assert.NoError(t, err)

	tx, err := db.BeginTxx(context.Background(), nil)
	assert.NoError(t, err)

	defer tx.Rollback()

	_, err = tx.Exec("drop table words")
	assert.NoError(t, err)

	_, err = tx.Exec("drop table users")
	assert.NoError(t, err)

	_, err = tx.Exec("create table users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)")
	assert.NoError(t, err)

	_, err = tx.Exec("create table words (id bigserial primary key,word varchar(55) not null,translate varchar(55) not null,created_at timestamp,user_id bigint references users(id))")
	assert.NoError(t, err)

	_, err = tx.Exec("insert into users (email, hash_password) values ($1, $2)", "bot@gmail.com", "qwerty")
	assert.NoError(t, err)

	_, err = tx.Exec("insert into words (word, translate, user_id) values ($1, $2, $3)", "boogaga", "смешняшка", 1)
	assert.NoError(t, err)

	wordDB := word.NewWordService(tx)

	req := httptest.NewRequest("DELETE", "/delete", bytes.NewBuffer([]byte(`{"word":"boogaga","user_id":1}`)))
	res := httptest.NewRecorder()

	handler := DeleteWordHandler(wordDB)
	handler(res, req)

	//err = wordDB.DeleteWordByWord("boogaga", 1)
	//assert.NoError(t, err)

}

func TestDeleteWordHandlerNegativeErrUnmarshal(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	assert.NoError(t, err)

	tx, err := db.BeginTxx(context.Background(), nil)
	assert.NoError(t, err)

	defer tx.Rollback()

	_, err = tx.Exec("drop table words")
	assert.NoError(t, err)

	_, err = tx.Exec("drop table users")
	assert.NoError(t, err)

	_, err = tx.Exec("create table users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)")
	assert.NoError(t, err)

	_, err = tx.Exec("create table words (id bigserial primary key,word varchar(55) not null,translate varchar(55) not null,created_at timestamp,user_id bigint references users(id))")
	assert.NoError(t, err)

	_, err = tx.Exec("insert into users (email, hash_password) values ($1, $2)", "bot@gmail.com", "qwerty")
	assert.NoError(t, err)

	_, err = tx.Exec("insert into words (word, translate, user_id) values ($1, $2, $3)", "boogaga", "смешняшка", 1)
	assert.NoError(t, err)

	wordDB := word.NewWordService(tx)

	req := httptest.NewRequest("DELETE", "/delete", nil)
	res := httptest.NewRecorder()

	handler := DeleteWordHandler(wordDB)
	handler(res, req)
}

func TestDeleteWordHandlerNegativeNotFound(t *testing.T) {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	assert.NoError(t, err)

	tx, err := db.BeginTxx(context.Background(), nil)
	assert.NoError(t, err)

	defer tx.Rollback()

	_, err = tx.Exec("drop table words")
	assert.NoError(t, err)

	_, err = tx.Exec("drop table users")
	assert.NoError(t, err)

	_, err = tx.Exec("create table users (id bigserial primary key,email varchar(55) unique not null ,hash_password varchar(255) not null)")
	assert.NoError(t, err)

	_, err = tx.Exec("create table words (id bigserial primary key,word varchar(55) not null,translate varchar(55) not null,created_at timestamp,user_id bigint references users(id))")
	assert.NoError(t, err)

	_, err = tx.Exec("insert into users (email, hash_password) values ($1, $2)", "bot@gmail.com", "qwerty")
	assert.NoError(t, err)

	_, err = tx.Exec("insert into words (word, translate, user_id) values ($1, $2, $3)", "boogaga", "смешняшка", 1)
	assert.NoError(t, err)

	wordDB := word.NewWordService(tx)

	err = wordDB.DeleteWordByWord("wrong word", 1)
	assert.ErrorIs(t, err, word.ErrWord)
}
