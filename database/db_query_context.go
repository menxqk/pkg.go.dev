package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Person struct {
	Name string
	Age  int
}

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
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	age := 45
	rows, err := db.QueryContext(ctx, "select p.name, p.age from people as p where age > $1 order by id asc", age)
	checkError(err)
	defer rows.Close()

	people := make([]Person, 0)
	for rows.Next() {
		var (
			name string
			age  int
		)
		if err := rows.Scan(&name, &age); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		people = append(people, Person{name, age})
	}
	// If the database is being written to ensure to check for Close
	// errors tha may be returned from the driver. The query may
	// encounter an auto-commit error and be forced to rollback changes.
	rerr := rows.Close()
	if rerr != nil {
		log.Fatal(rerr)
	}

	// Rows.Err will report tha last error encountered by Rows.Scan.
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	for _, person := range people {
		fmt.Printf("%s is %d years old.\n", person.Name, person.Age)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
