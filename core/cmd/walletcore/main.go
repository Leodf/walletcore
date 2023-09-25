package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com.br/Leodf/walletcore/internal/database"
	"github.com.br/Leodf/walletcore/internal/event"
	"github.com.br/Leodf/walletcore/internal/event/handler"
	createAccount "github.com.br/Leodf/walletcore/internal/usecase/create-account"
	createClient "github.com.br/Leodf/walletcore/internal/usecase/create-client"
	createTransaction "github.com.br/Leodf/walletcore/internal/usecase/create-transaction"
	"github.com.br/Leodf/walletcore/internal/web"
	"github.com.br/Leodf/walletcore/internal/web/webserver"
	"github.com.br/Leodf/walletcore/pkg/events"
	"github.com.br/Leodf/walletcore/pkg/kafka"
	unitofwork "github.com.br/Leodf/walletcore/pkg/unit-of-work"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "mysql", "3306", "wallet"))
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database is connected")
	defer db.Close()

	configMap := ckafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		"group.id":          "wallet",
	}
	kafkaProducer := kafka.NewKafkaProducer(&configMap)

	eventDispatcher := events.NewEventDispatcher()
	transactionCreatedEvent := event.NewTransactionCreated()
	balanceUpdatedEvent := event.NewBalanceUpdated()
	eventDispatcher.Register("TransactionCreated", handler.NewTransactionCreatedKafkaHandler(kafkaProducer))
	eventDispatcher.Register("BalanceUpdated", handler.NewBalanceUpdatedKafkaHandler(kafkaProducer))

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
	createTransactionUseCase := createTransaction.NewCreateTransactionUseCase(uow, eventDispatcher, transactionCreatedEvent, balanceUpdatedEvent)

	webServer := webserver.NewWebServer(":8080")

	clientHandler := web.NewWebClientHandler(*createClientUseCase)
	accountHandler := web.NewWebAccountHandler(*createAccountUseCase)
	transactionHandler := web.NewWebTransactionHandler(*createTransactionUseCase)

	webServer.AddHandler("/clients", clientHandler.CreateClient)
	webServer.AddHandler("/accounts", accountHandler.CreateAccount)
	webServer.AddHandler("/transactions", transactionHandler.CreateTransaction)

	fmt.Println("Server is running")
	webServer.Start()
}
