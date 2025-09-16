package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Conn() (*sql.DB, error) {
	if err := godotenv.Load("../../.env"); err != nil {
		return nil, err
	}
	log.Println(os.Getenv("DBUSER"))
	log.Println(os.Getenv("DBPASS"))
	log.Println(os.Getenv("DBNAME"))

	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = os.Getenv("DBNAME")

	// Get a database handle.
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, errors.New("Falha ao abrir conexão!")
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
	return db, nil
}
