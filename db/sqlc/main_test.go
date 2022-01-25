package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var conn *sql.DB
const (
	dbDriver = "postgres"
	dbSource = "postgresql://nothing:nothing@localhost:5432/bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	conn, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalln(err)
	}
	testQueries = New(conn)
	os.Exit(m.Run())
}
