package db

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func ConnectPostgres() *sql.DB{
	connStr := "user=postgres password=test dbname=hospital sslmode=disable"
	db,err := sql.Open("postgres", connStr)
	if err != nil{
		log.Fatal("Error connecting to postgres Db :", err)
	}
	return db
}
