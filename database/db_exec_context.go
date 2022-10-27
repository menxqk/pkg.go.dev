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
		log.Fatal(err)
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := db.ExecContext(ctx, "update people set active=false;")
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	if rows != 1 {
		log.Fatalf("expected to affect 1 row, affected %d", rows)
	}
}
