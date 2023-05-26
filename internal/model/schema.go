package model

import (
	"drunklish/internal/pkg/db"
)

const (
	createUserTable = `create table if not exists users (
    id bigserial primary key,
    email varchar(55) unique not null ,
    hash_password varchar(255) not null
);`
	createWordsTable = `create table if not exists words (
    id bigserial primary key,
    word varchar(55) not null,
    translate varchar(55) not null,
    created_at timestamp,
    user_id bigint not null references users(id) 
);`
	createTrainingTable = `create table if not exists training (
    id bigserial primary key,
    words jsonb not null,
    answers jsonb not null,
    words_total bigint not null,
    user_id bigint references users(id)
);`
	createStatisticTable = `create table if not exists statistic (
    id bigserial primary key,
    training_id bigint references training(id),
    correct_answers jsonb,
    wrong_answers jsonb,
    user_id bigint references users(id)
);`
)

func CreateTables(s *db.Storage) error {
	tx := s.TX.MustBegin()
	tx.MustExec(createUserTable)
	tx.MustExec(createWordsTable)
	tx.MustExec(createTrainingTable)
	tx.MustExec(createStatisticTable)

	return tx.Commit()
}
