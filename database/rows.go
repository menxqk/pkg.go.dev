package main

import (
	"context"
	"database/sql"
	"fmt"
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

	rows, err := db.QueryContext(ctx, "select name from people")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	names := make([]string, 0)
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}

	// Check for errors from iterating over rows.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, name := range names {
		fmt.Printf("%s\n", name)
	}
}
