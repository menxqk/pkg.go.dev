package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost dbname=test user=test password=test sslmode=disable")
	if err != nil {
		// This will not be connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	db.SetConnMaxIdleTime(0)
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(50)

	s := &Service{db: db}
	http.ListenAndServe(":8080", s)
}

type Service struct {
	db *sql.DB
}

func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	db := s.db
	switch r.URL.Path {
	default:
		http.Error(w, "not found", http.StatusNotFound)
		return
	case "/healthz":
		ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
		defer cancel()

		err := s.db.PingContext(ctx)
		if err != nil {
			http.Error(w, fmt.Sprintf("db down: %v", err), http.StatusFailedDependency)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	case "/quick-action":
		// This is a short SELECT. Use the request context as the base of
		// the context timeout.
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		id := 5
		var name string
		err := db.QueryRowContext(ctx, "select p.name from people as p where p.id = $1;", id).Scan(&name)
		if err != nil {
			if err == sql.ErrNoRows {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		io.WriteString(w, name)
		return
	case "/long-action":
		// This is a long SELECT. Use the request context as the base of
		// the context timeout, but give it some time to finish. If
		// the client cancels before the query is done the query will also
		// be canceled.
		ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
		defer cancel()

		var names []string
		rows, err := db.QueryContext(ctx, "select p.name from people as p where p.active = true;")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for rows.Next() {
			var name string
			err = rows.Scan(&name)
			if err != nil {
				break
			}
			names = append(names, name)
		}
		// Check for errors during rows "Close".
		// This may be more important if multiple statements are executed
		// in a single batch and rows were written as well read.
		if closeErr := rows.Close(); closeErr != nil {
			http.Error(w, closeErr.Error(), http.StatusInternalServerError)
			return
		}

		// Check for row scan error.
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check for errors during row iteration.
		if err = rows.Err(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(names)
		return

	case "/async-action":
		// This action has side effects that we want to preserve
		// even if the client cancels the HTTP request part way through.
		// For this we do not use the http request context as a base for the
		// timeout.
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		var id = 2
		tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
		_, err = tx.ExecContext(ctx, "update people set active=true where id=$1", id)
		if err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = tx.Commit()
		if err != nil {
			http.Error(w, "action in unknown state, check state before attempting again", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
