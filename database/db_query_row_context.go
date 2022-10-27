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
	checkError(err)

	err = db.Ping()
	checkError(err)

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	id := 3
	var name string
	err = db.QueryRowContext(ctx, "select p.name from people as p where p.id=$1;", id).Scan(&name)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("no user with id %d\n", id)
	case err != nil:
		log.Fatalf("query error: %v\n", err)
	default:
		log.Printf("username is %q with id=%d\n", name, id)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
