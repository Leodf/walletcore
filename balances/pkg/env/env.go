package env

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitEnvFiles() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Coudn't load .env file")
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		err := fmt.Errorf("missing environment variable %s", key)
		panic(err)
	}

	return value
}
