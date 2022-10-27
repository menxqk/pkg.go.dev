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

	// In normal use, create one stmt when your process starts.
	stmt, err := db.Prepare("select p.name from people as p where id = $1;")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Then reuse it each time you need to issue the query.
	id := 4
	var name string
	err = stmt.QueryRowContext(ctx, id).Scan(&name)
	switch {
	case err == sql.ErrNoRows:
		log.Fatalf("no user with id %d", id)
	case err != nil:
		log.Fatal(err)
	default:
		log.Printf("user %q has id %d\n", name, id)
	}
}
