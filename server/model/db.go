package model

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func SetUp() {
	_db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		// TODO: slogでかく
		log.Fatalf("Cannot Connect to Database: %s", err)
	}
	db = _db
	// TODO: slogでかく
	log.Printf("Connected to Database")
}
