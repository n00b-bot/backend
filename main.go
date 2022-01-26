package main

import (
	"backend/api"
	db "backend/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://nothing:nothing@localhost:5432/bank?sslmode=disable"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln(err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	server.Start(":9090")
}
