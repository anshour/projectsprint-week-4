package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var isEnvLoaded = false

var (
	DB_NAME     = getEnv("DB_NAME", "")
	DB_PORT     = getEnv("DB_PORT", "")
	DB_HOST     = getEnv("DB_HOST", "")
	DB_USERNAME = getEnv("DB_USERNAME", "")
	DB_PASSWORD = getEnv("DB_PASSWORD", "")
	DB_PARAMS   = getEnv("DB_PARAMS", "sslmode=disable")

	JWT_SECRET  = getEnv("JWT_SECRET", "")
	JWT_EXP     = getEnv("JWT_EXP", "8h")
	BCRYPT_SALT = getEnvAsInt("BCRYPT_SALT", 8)

	AWS_BUCKET            = getEnv("AWS_S3_BUCKET_NAME", "")
	AWS_REGION            = getEnv("AWS_REGION", "")
	AWS_ACCESS_KEY_ID     = getEnv("AWS_ACCESS_KEY_ID", "")
	AWS_SECRET_ACCESS_KEY = getEnv("AWS_SECRET_ACCESS_KEY", "")
)

func loadEnv() {
	if !isEnvLoaded {
		err := godotenv.Load()
		if err != nil {
			fmt.Println(".env file is not found, using env from os...")
		}
		isEnvLoaded = true
	}

}

func getEnv(name string, fallback string) string {
	loadEnv()
	if value := os.Getenv(name); value != "" {
		return value
	}

	if fallback != "" {
		return fallback
	}

	panic(fmt.Sprintf(`Environment variable not found :: %v`, name))
}

func getEnvAsInt(name string, fallback int) int {
	loadEnv()
	if value := os.Getenv(name); value != "" {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatalf("Error converting %v to int", name)
		}
		return intValue
	}

	return fallback
}
