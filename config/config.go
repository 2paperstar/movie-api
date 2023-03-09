package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DbUrl     string
	JwtSecret []byte
)

func init() {
	godotenv.Load()
	DbUrl = os.Getenv("DB_URL")
	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
}
