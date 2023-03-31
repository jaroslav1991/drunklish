package main

import (
	"drunklish/internal/api"
	"drunklish/internal/config"
	"drunklish/internal/connection"
	"drunklish/internal/model"
	"drunklish/internal/service"
	"drunklish/internal/service/auth"
	authRepo "drunklish/internal/service/auth/repository"
	"drunklish/internal/service/word"
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

	storageDB := service.NewStorage(db, tx)
	authDB := auth.NewAuthService(db, authRepo.NewAuthRepository(db))
	wordDB := word.NewWordService(db, wordRepo.NewWordRepository(db))

	if err := model.CreateTables(storageDB); err != nil {
		log.Fatal(err)
	}

	http.Handle("/sign-up", api.SignUpHandler(authDB))
	http.Handle("/sign-in", api.SignInHandler(authDB))
	http.Handle("/word", api.CreateWordHandler(wordDB))
	http.Handle("/get-words", api.GetWordsHandler(wordDB))
	http.Handle("/delete", api.DeleteWordHandler(wordDB))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
