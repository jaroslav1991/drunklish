package config

import "os"

var localdbConfig = "postgresql://postgres:1234@localhost:5432/drunklish?sslmode=disable"

func GetDBConfig() string {
	if s := os.Getenv("PG_DSN"); s != "" {
		return s
	}
	return localdbConfig
}
