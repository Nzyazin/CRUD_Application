package main

import (
	"database/sql"
	"fmt"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "43dagtot21"
	dbname   = "library"
)

func main() {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
	insert := "inseret into 'author('author_id', 'author_name', 'birthday') values
}
