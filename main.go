package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "43dagtot21"
	dbname   = "library"
)

func main() {
	psqlconn := fmt.Sprintf("host= %s port = %d user = %s password =%s dbname = %s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Could not connect", err)
	}
	defer db.Close()
	var (
		authorId int
		name     string
		birthDay int
	)
	rows, err := db.Query("INSERT INTO author(author_id, author_name, birthday) values ('10', 'Petrov K. P.', 'birthday')", 1)
	for rows.Next() {
		err := rows.Scan(&authorId, &name, &birthDay)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(authorId, name, birthDay)
	}
	if err != nil {
		fmt.Println("Can not:", err)
	}
	defer rows.Close()
}
