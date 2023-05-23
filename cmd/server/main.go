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

	http.Handle("/api/sign-up", httputils.WrapRpc(authHandlers.SignUpHandler(authDB)))
	http.Handle("/api/sign-in", httputils.WrapRpc(authHandlers.SignInHandler(authDB)))
	http.Handle("/api/word", httputils.WrapRpc(wordsHandlers.CreateWordHandler(wordDB)))
	http.Handle("/api/get-words", httputils.WrapRpc(wordsHandlers.GetWordsHandler(wordDB)))
	http.Handle("/api/get-words-period", httputils.WrapRpc(wordsHandlers.GetWordByPeriodHandler(wordDB)))
	http.Handle("/api/delete", httputils.WrapRpc(wordsHandlers.DeleteWordHandler(wordDB)))
	http.Handle("/api/update", httputils.WrapRpc(wordsHandlers.UpdateWordHandler(wordDB)))
	http.Handle("/api/training", httputils.WrapRpc(wordsHandlers.CreateTrainingHandler(wordDB)))
	http.Handle("/api/get-statistic", httputils.WrapRpc(wordsHandlers.GetStatisticHandler(wordDB)))
	http.Handle("/api/statistic", httputils.WrapRpc(wordsHandlers.CreateStatisticHandler(wordDB)))

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
