package utils

import "os"

var (
	Environment map[string]string
)

func InitEnvironment() {
	Environment = map[string]string{
		"DISCORD_TOKEN":     os.Getenv("DISCORD_TOKEN"),
		"JWT_PASSWORD":      os.Getenv("JWT_PASSWORD"),
		"POSTGRES_HOST":     os.Getenv("POSTGRES_HOST"),
		"POSTGRES_PORT":     os.Getenv("POSTGRES_PORT"),
		"POSTGRES_USER":     os.Getenv("POSTGRES_USER"),
		"POSTGRES_PASSWORD": os.Getenv("POSTGRES_PASSWORD"),
		"DB_NAME":           os.Getenv("DB_NAME"),
		"PORT_API":          os.Getenv("PORT_API"),
		"PORT_WEBSITE":      os.Getenv("PORT_WEBSITE"),
	}
}
