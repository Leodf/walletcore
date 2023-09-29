package connection

import (
	"database/sql"
	"fmt"

	"github.com.br/Leodf/walletcore/balances/pkg/env"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDatabase() *sql.DB {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			env.GetEnv("DB_USER"),
			env.GetEnv("DB_PASSWORD"),
			env.GetEnv("DB_HOST"),
			env.GetEnv("DB_PORT"),
			env.GetEnv("DB_NAME"),
		))
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database is connected")
	Db = db
	return db
}
