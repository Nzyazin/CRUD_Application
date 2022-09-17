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

func create(a_id int, a_name string, b_day string) {
	psqlconn := fmt.Sprintf("host= %s port = %d user = %s password =%s dbname = %s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Could not connect", err)
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO author(author_id, author_name, birthday) VALUES($1, $2, $3)")
	if err != nil {
		log.Fatal("Here2: ", err)
	}
	res, err := stmt.Exec(a_id, a_name, b_day)
	if err != nil {
		log.Fatal("Here1: ", err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n Succeded!!!", res, rowCnt)
}

func main() {
	variable := true
	for variable {
		var operation int
		fmt.Println("Input a number for procedure with database or for quit: \n1. Create a field in table. \n2. Read a table. \n3. Update a field in table. \n4. Delete a field. \n5. Quit")
		fmt.Scanln(&operation)
		switch operation {
		case 1:
			fmt.Println("Input author id: ")
			var authorId int
			fmt.Scanln(&authorId)
			fmt.Println("Input author author name:")
			var name string
			fmt.Scanln(&name)
			fmt.Println("And his birthday:")
			var birthday string
			fmt.Scanln(&birthday)
			create(authorId, name, birthday)
		case 5:
			variable = false
			fmt.Println("Bye bye")
		}
	}
	psqlconn := fmt.Sprintf("host= %s port = %d user = %s password =%s dbname = %s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Could not connect", err)
	}
	defer db.Close()

	var (
		authorId int
		name     sql.NullString
		birthday sql.NullString
	)
	rows, err := db.Query("select * from author")
	if err != nil {
		log.Fatal("Did`not query:", err)
	}

	for rows.Next() {
		err := rows.Scan(&authorId, &name, &birthday)
		if err != nil {
			log.Fatal("Did`not connect", err)
		}
		log.Println(authorId, name.String, birthday.String)
	}
	if err != nil {
		fmt.Println("Can not:", err)
	}
	defer rows.Close()
}
