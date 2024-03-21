package database

import (
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Erro ao iniciar as variáveis de ambiente")
	}

	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    "tcp",
		Addr:   os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT"),
		DBName: os.Getenv("DB_NAME"),
	}

	var openErr error
	db, openErr = sql.Open("mysql", cfg.FormatDSN())
	if openErr != nil {
		panic(openErr)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}
}
