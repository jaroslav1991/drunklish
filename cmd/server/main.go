package main

import (
	"drunklish/internal/config"
	"drunklish/internal/connection"
	"drunklish/internal/model"
	pkgdb "drunklish/internal/pkg/db"
	"drunklish/internal/pkg/httputils"
	"drunklish/internal/service/auth"
	authHandlers "drunklish/internal/service/auth/handlers"
	authRepo "drunklish/internal/service/auth/repository"
	"drunklish/internal/service/word"
	wordsHandlers "drunklish/internal/service/word/handlers"
	wordRepo "drunklish/internal/service/word/repository"
	"log"
	"net/http"
)

func main() {
	dbConfig := config.GetDBConfig()
	db, err := connection.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := connection.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	storageDB := pkgdb.NewStorage(db, tx)
	authDB := auth.NewAuthService(authRepo.NewAuthRepository(db))
	wordDB := word.NewWordService(wordRepo.NewWordRepository(db))

	if err := model.CreateTables(storageDB); err != nil {
		log.Fatal(err)
	}

	http.Handle("/sign-up", httputils.WrapRpc(authHandlers.SignUpHandler(authDB)))
	http.Handle("/sign-in", httputils.WrapRpc(authHandlers.SignInHandler(authDB)))
	http.Handle("/word", wordsHandlers.CreateWordHandler(wordDB))
	http.Handle("/get-words", wordsHandlers.GetWordsHandler(wordDB))
	http.Handle("/delete", wordsHandlers.DeleteWordHandler(wordDB))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
