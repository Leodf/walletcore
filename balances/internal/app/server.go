package app

import (
	"log"

	"github.com.br/Leodf/walletcore/balances/internal/app/config"
	"github.com.br/Leodf/walletcore/balances/pkg/env"
)

func Run() {
	app := config.InitApp()
	err := app.Listen(env.GetEnv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
