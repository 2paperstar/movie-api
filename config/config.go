package config

import (
	"os"

	"github.com/joho/godotenv"
)

var DbUrl string

func init() {
	godotenv.Load()
	DbUrl = os.Getenv("DB_URL")
}
