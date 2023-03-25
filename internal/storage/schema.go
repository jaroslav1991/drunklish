package storage

const (
	createUserTable = `create table if not exists users (
    id bigserial primary key,
    nick_name varchar(55) unique not null 
);`
	createWordsTable = `create table if not exists words (
    id bigserial primary key,
    word varchar(55) not null,
    translate varchar(55) not null,
    user_id bigint references users(id) 
);`
)

func CreateTables(s *Storage) error {
	_, err := s.DB.Exec(createUserTable)
	_, err = s.DB.Exec(createWordsTable)
	return err
}
