package database

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "password"
	DB       *sql.DB
)

func Initialize() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s  sslmode=disable",
		host, port, user, password)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB CONNECTED")

	return err
}

func CheckDB() bool {
	var (
		err    error
		exists bool
	)

	err = DB.QueryRow("SELECT EXISTS(SELECT * FROM information_schema.tables WHERE table_schema = $1 AND table_name = $2);", "public", "words").Scan(&exists)

	if err != nil {
		log.Fatal(err)
	}

	return exists
}
