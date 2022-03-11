package database

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func connect() {
	fmt.Println("Connecting to database...")

	var connStr string = "username:password@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", connStr)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
