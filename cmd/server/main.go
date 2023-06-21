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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"os"
)

func init() {
	fileInfo, err := os.OpenFile("drunklish-logger", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Println(err)
		return
	}
	log.SetOutput(fileInfo)

}

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

	reg := prometheus.NewRegistry()

	metrics := httputils.NewMetrics(reg)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	http.Handle("/api/sign-in", metrics.Middleware(metrics, httputils.WrapRpc(metrics, authHandlers.SignInHandler(authDB))))
	http.Handle("/api/sign-up", metrics.Middleware(metrics, httputils.WrapRpc(metrics, authHandlers.SignUpHandler(authDB))))
	http.Handle("/api/word", metrics.Middleware(metrics, httputils.WrapRpc(metrics, wordsHandlers.CreateWordHandler(wordDB))))
	http.Handle("/api/get-words", metrics.Middleware(metrics, httputils.WrapRpc(metrics, wordsHandlers.GetWordsHandler(wordDB))))
	http.Handle("/api/get-words-period", metrics.Middleware(metrics, httputils.WrapRpc(metrics, wordsHandlers.GetWordByPeriodHandler(wordDB))))
	http.Handle("/api/delete", metrics.Middleware(metrics, httputils.WrapRpc(metrics, wordsHandlers.DeleteWordHandler(wordDB))))
	http.Handle("/api/update", metrics.Middleware(metrics, httputils.WrapRpc(metrics, wordsHandlers.UpdateWordHandler(wordDB))))
	http.Handle("/api/training", metrics.Middleware(metrics, httputils.WrapRpc(metrics, wordsHandlers.CreateTrainingHandler(wordDB))))
	http.Handle("/api/get-statistic", metrics.Middleware(metrics, httputils.WrapRpc(metrics, wordsHandlers.GetStatisticHandler(wordDB))))
	http.Handle("/api/statistic", metrics.Middleware(metrics, httputils.WrapRpc(metrics, wordsHandlers.CreateStatisticHandler(wordDB))))

	if err := http.ListenAndServe(":8585", nil); err != nil {
		log.Fatal(err)
	}

}
