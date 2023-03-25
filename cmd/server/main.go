package main

import (
	"drunklish/internal/config"
	"drunklish/internal/storage"
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

	storageDB := storage.NewStorage(db)

	if err := storage.CreateTables(storageDB); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
