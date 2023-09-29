package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	GetEnv(key string) string
}

type configImpl struct {
}

func (config *configImpl) GetEnv(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	if err != nil {
		panic(err)
	}
	return &configImpl{}
}
