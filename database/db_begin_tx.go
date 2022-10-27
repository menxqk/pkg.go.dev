package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var (
	ctx context.Context
	db  *sql.DB
)

func main() {
	var err error
	db, err = sql.Open("postgres", "host=localhost dbname=test user=test password=test sslmode=disable")
	if err != nil {
		// This will not be connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxIdleTime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}

	id := 5
	_, execErr := tx.Exec("update people set active=true where id=$1;", id)
	if execErr != nil {
		_ = tx.Rollback()
		log.Fatal(execErr)
	}
	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
