package storage

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
    user_id bigint references users(id) 
);`
)

func CreateTables(s *Storage) error {
	tx := s.DB.MustBegin()
	tx.MustExec(createUserTable)
	tx.MustExec(createWordsTable)

	return tx.Commit()
}
