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
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	// The rollback will be ignored if the tx has been committed
	// later in the function.

	stmt, err := tx.PrepareContext(ctx, "update people set active=true where id = $1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	ids := []int{1, 3, 5}
	for _, id := range ids {
		if _, err := stmt.ExecContext(ctx, id); err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				log.Fatalf("could not rollback: %v\n", rollbackErr)
			}
			log.Fatal(err)
		}
	}

	if err := tx.Commit(); err != nil {
		log.Fatal(err)
	}
}
