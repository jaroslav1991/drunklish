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
	"go.uber.org/zap"
	"log"
	"net/http"
)

func main() {

	logger, err := httputils.ConfigLogger("drunklish-log")
	if err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	dbConfig := config.GetDBConfig()
	db, err := connection.NewPostgresDB(dbConfig)
	if err != nil {
		logger.Fatal("can't connect to postgres db", zap.Error(err))
	}

	tx, err := connection.NewPostgresDB(dbConfig)
	if err != nil {
		logger.Fatal("can't connect to postgres tx", zap.Error(err))
	}

	storageDB := pkgdb.NewStorage(db, tx)
	authDB := auth.NewAuthService(authRepo.NewAuthRepository(db))
	wordDB := word.NewWordService(wordRepo.NewWordRepository(db))

	if err := model.CreateTables(storageDB); err != nil {
		logger.Fatal("can't create tables in postgres", zap.Error(err))
	}

	reg := prometheus.NewRegistry()

	metrics := httputils.NewMetrics(reg)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	http.Handle("/api/sign-in", metrics.Wrap(httputils.WrapRpc(logger, metrics, authHandlers.SignInHandler(authDB))))
	http.Handle("/api/sign-up", metrics.Wrap(httputils.WrapRpc(logger, metrics, authHandlers.SignUpHandler(authDB))))
	http.Handle("/api/word", metrics.Wrap(httputils.WrapRpc(logger, metrics, wordsHandlers.CreateWordHandler(wordDB))))
	http.Handle("/api/get-words", metrics.Wrap(httputils.WrapRpc(logger, metrics, wordsHandlers.GetWordsHandler(wordDB))))
	http.Handle("/api/get-words-period", metrics.Wrap(httputils.WrapRpc(logger, metrics, wordsHandlers.GetWordByPeriodHandler(wordDB))))
	http.Handle("/api/delete", metrics.Wrap(httputils.WrapRpc(logger, metrics, wordsHandlers.DeleteWordHandler(wordDB))))
	http.Handle("/api/update", metrics.Wrap(httputils.WrapRpc(logger, metrics, wordsHandlers.UpdateWordHandler(wordDB))))
	http.Handle("/api/training", metrics.Wrap(httputils.WrapRpc(logger, metrics, wordsHandlers.CreateTrainingHandler(wordDB))))
	http.Handle("/api/get-statistic", metrics.Wrap(httputils.WrapRpc(logger, metrics, wordsHandlers.GetStatisticHandler(wordDB))))
	http.Handle("/api/statistic", metrics.Wrap(httputils.WrapRpc(logger, metrics, wordsHandlers.CreateStatisticHandler(wordDB))))

	if err := http.ListenAndServe(":8585", nil); err != nil {
		logger.Fatal("failed with listen and serve", zap.Error(err))
	}

}
