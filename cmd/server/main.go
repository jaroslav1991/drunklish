package main

import (
	"drunklish/internal/config"
	"drunklish/internal/connection"
	"drunklish/internal/model"
	"drunklish/internal/service"
	"drunklish/internal/service/auth"
	"drunklish/internal/service/auth/handlers"
	authRepo "drunklish/internal/service/auth/repository"
	"drunklish/internal/service/word"
	handlers2 "drunklish/internal/service/word/handlers"
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
	authDB := auth.NewAuthService(authRepo.NewAuthRepository(db))
	wordDB := word.NewWordService(wordRepo.NewWordRepository(db))

	if err := model.CreateTables(storageDB); err != nil {
		log.Fatal(err)
	}

	http.Handle("/sign-up", handlers.SignUpHandler(authDB))
	http.Handle("/sign-in", handlers.SignInHandler(authDB))
	http.Handle("/word", handlers2.CreateWordHandler(wordDB))
	http.Handle("/get-words", handlers2.GetWordsHandler(wordDB))
	http.Handle("/delete", handlers2.DeleteWordHandler(wordDB))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
