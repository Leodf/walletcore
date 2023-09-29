package env

import (
	"fmt"
	"os"
)

func GetEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		err := fmt.Errorf("missing environment variable %s", key)
		panic(err)
	}

	return value
}
