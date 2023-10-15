package app

import (
	"log"

	"github.com.br/Leodf/walletcore/balances/internal/app/config"

	"github.com.br/Leodf/walletcore/balances/internal/infra/database/connection"
	"github.com.br/Leodf/walletcore/balances/internal/infra/env"
	"github.com.br/Leodf/walletcore/balances/internal/infra/kafka"
)

func Run() {
	connection.InitDatabase()
	defer connection.DB.Close()
	app := config.InitApp()
	kafka.RunKafka()
	err := app.Listen(env.GetEnv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
