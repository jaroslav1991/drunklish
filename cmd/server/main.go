package main

import (
	"drunklish/internal/api"
	"drunklish/internal/config"
	"drunklish/internal/model"
	"drunklish/internal/service"
	"drunklish/internal/service/auth"
	"drunklish/internal/service/word"
	"drunklish/pkg/repository"
	"log"
	"net/http"
)

func main() {
	dbConfig := config.GetDBConfig()
	db, err := repository.NewPostgresDB(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	storageDB := service.NewStorage(db)
	authDB := auth.NewAuthService(db)
	wordDB := word.NewWordService(db)

	if err := model.CreateTables(storageDB); err != nil {
		log.Fatal(err)
	}

	http.Handle("/sign-up", api.SignUpHandler(authDB))
	http.Handle("/sign-in", api.SignInHandler(authDB))
	http.Handle("/word", api.CreateWordHandler(wordDB))
	http.Handle("/get-words", api.GetWordsHandler(wordDB))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
