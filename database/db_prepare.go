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
	ctx   context.Context
	db    *sql.DB
	names = []string{"Larry", "Gonzalo", "Alex", "Alexander", "Brian"}
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

	ping := pingDb(ctx, db)
	if !ping {
		log.Fatalf("could not ping database")
	}

	err = initDb(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	err = populateDb(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	err = updateDb(ctx, db)
	if err != nil {
		log.Fatal(err)
	}
}

func updateDb(ctx context.Context, db *sql.DB) error {
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return (err)
	}
	_, execErr := tx.Exec("update people set active=true where id in (1, 3, 5);")
	if execErr != nil {
		_ = tx.Rollback()
		return execErr
	}
	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func populateDb(ctx context.Context, db *sql.DB) error {
	stmt, err := db.PrepareContext(ctx, "insert into people (name) values ($1);")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()

	for _, name := range names {
		_, err := stmt.ExecContext(ctx, name)
		if err != nil {
			return err
		}
	}
	return nil
}

func initDb(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, "drop table if exists people;")
	if err != nil {
		return err
	}
	_, err = db.ExecContext(ctx, "create table people (id serial primary key, name varchar(60), active boolean default false);")
	if err != nil {
		return err
	}
	return nil
}

func pingDb(ctx context.Context, db *sql.DB) bool {
	err := db.PingContext(ctx)
	if err == nil {
		return true
	}
	return false
}
