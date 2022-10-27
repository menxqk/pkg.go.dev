package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("postgres", "host=localhost dbname=test user=test password=test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	age := 50

	q := `
-- First result set.
select p.id, p.name from people as p where p.age >= $1 order by id asc;	
`
	stmt, err := db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmt.Query(age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Printf("Age >= %d: id %d name is %s\n", age, id, name)
	}

	q = `
-- Second result set.
select p.id, p.name from people as p where p.age < $1 order by id asc;
`
	stmt, err = db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}

	rows, err = stmt.Query(age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Printf("Age < %d: id %d name is %s\n", age, id, name)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
