package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase(config Config) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.GetEnv("DB_USER"),
			config.GetEnv("DB_PASSWORD"),
			config.GetEnv("DB_HOST"),
			config.GetEnv("DB_PORT"),
			config.GetEnv("DB_NAME"),
		))
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Database is connected")
	DB = db
}
