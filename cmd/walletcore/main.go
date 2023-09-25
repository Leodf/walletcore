package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com.br/Leodf/walletcore/internal/database"
	"github.com.br/Leodf/walletcore/internal/event"
	createAccount "github.com.br/Leodf/walletcore/internal/usecase/create-account"
	createClient "github.com.br/Leodf/walletcore/internal/usecase/create-client"
	createTransaction "github.com.br/Leodf/walletcore/internal/usecase/create-transaction"
	"github.com.br/Leodf/walletcore/internal/web"
	"github.com.br/Leodf/walletcore/internal/web/webserver"
	"github.com.br/Leodf/walletcore/pkg/events"
	unitofwork "github.com.br/Leodf/walletcore/pkg/unit-of-work"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "localhost", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database is connected")
	defer db.Close()

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	// eventDispatcher.Register("TransactionCreated", handler)
	clientDb := database.NewClientDB(db)
	accountDb := database.NewAccountDB(db)

	ctx := context.Background()
	uow := unitofwork.NewUow(ctx, db)

	uow.Register("AccountDB", func(tx *sql.Tx) interface{} {
		return database.NewAccountDB(db)
	})
	uow.Register("TransactionDB", func(tx *sql.Tx) interface{} {
		return database.NewTransactionDB(db)
	})

	createClientUseCase := createClient.NewCreateClientUseCase(clientDb)
	createAccountUseCase := createAccount.NewCreateAccountUseCase(accountDb, clientDb)
	createTransactionUseCase := createTransaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent)

	webServer := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webServer.AddHandler("/clients", clientHandler.CreateClient)
	webServer.AddHandler("/accounts", accountHandler.CreateAccount)
	webServer.AddHandler("/transactions", transactionHandler.CreateTransaction)

	webServer.Start()
}
