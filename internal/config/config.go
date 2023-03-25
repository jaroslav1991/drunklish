package config

import "os"

var localdbConfig = "postgresql://postgres:1234@localhost:5432/postgres?application_name=drunklish&sslmode=disable"

func GetDBConfig() string {
	if s := os.Getenv("PG_DSN"); s != "" {
		return s
	}
	return localdbConfig
}
