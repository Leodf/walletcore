package main

import (
	"github.com.br/Leodf/walletcore/balances/internal/app"
	"github.com.br/Leodf/walletcore/balances/internal/database/connection"
	"github.com.br/Leodf/walletcore/balances/pkg/env"
)

func main() {
	env.InitEnvFiles()
	db := connection.InitDatabase()
	defer db.Close()
	app.Run()
}
