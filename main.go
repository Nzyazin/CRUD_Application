package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "43dagtot21"
	dbname   = "library"
)

func create(a_name string, b_day string) {
	psqlconn := fmt.Sprintf("host= %s port = %d user = %s password =%s dbname = %s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		fmt.Println("Could not connect", err)
	}
	defer db.Close()
	var author_id int
	rows, err := db.Query("SELECT author_id from author order by author_id desc limit 1")
	for rows.Next() {
		err1 := rows.Scan(&author_id)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
	stmt, err := db.Prepare("INSERT INTO author(author_id, author_name, birthday) VALUES($1, $2, $3)")
	if err != nil {
		log.Fatal("Here1: ", err)
	}
	res, err := stmt.Exec(author_id+1, a_name, b_day)
	if err != nil {
		log.Fatal("Here2: ", err)
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
		fmt.Println("Input a number for procedure with database or for quit: \n1. Add a field in table. \n2. Output all existing table in database. \n3. Read a table. \n4. Update a field in table. \n5. Delete a field. \n6. Quit")
		fmt.Scanln(&operation)
		psqlconn := fmt.Sprintf("host= %s port = %d user = %s password =%s dbname = %s sslmode=disable", host, port, user, password, dbname)
		switch operation {
		case 1:
			fmt.Println("Input author name:")
			var name string
			in := bufio.NewReader(os.Stdin)
			name, err := in.ReadString('\n')
			if err != nil {
				log.Println(err)
			}
			fmt.Println("And his birthday in format <YYYY-MM-DD>:")
			var birthday string
			fmt.Scanln(&birthday)

			create(name, birthday)
		case 2:
			db, err := sql.Open("postgres", psqlconn)
			if err != nil {
				fmt.Println("Could not connect", err)
			}
			defer db.Close()
			rows, err := db.Query("SELECT tablename\nFROM pg_catalog.pg_tables\nWHERE schemaname != 'pg_catalog'\nAND schemaname != 'information_schema'\nORDER BY tablename ASC;")
			if err != nil {
				log.Fatal("Did`not query:", err)
			}
			var tables_name string
			for rows.Next() {
				err := rows.Scan(&tables_name)
				if err != nil {
					log.Fatal("Did`not connect", err)
				}
				log.Println(tables_name)
			}
		case 3:
			fmt.Println("Input name table (if you don`t know, you can use 2nd option): ")

			db, err := sql.Open("postgres", psqlconn)
			if err != nil {
				fmt.Println("Could not connect", err)
			}
			defer db.Close()
			var (
				author_id int
				name      sql.NullString
				birthday  sql.NullString
			)
			in := bufio.NewReader(os.Stdin)
			name_table, err := in.ReadString('\n')
			if err != nil {
				log.Println(err)
			}
			name_table = name_table[:len(name_table)-2]
			switch name_table {
			case "author":
				rows, err := db.Query("select * from author")
				if err != nil {
					log.Fatal("Did`not query:", err)
				}

				for rows.Next() {
					err1 := rows.Scan(&author_id, &name, &birthday)
					if err1 != nil {
						log.Fatal("Did`not scan: ", err1)
					}
					log.Println(author_id, name.String, birthday.String)
				}
				if err != nil {
					fmt.Println("Can not:", err)
				}
				defer rows.Close()
			case "books":
				var (
					id          int
					book_name   sql.NullString
					author_name sql.NullString
					publisher   sql.NullString
				)
				rows, err := db.Query("select books.book_id, books.book_name, author.author_name as author, publishing_house.ph_name as publisher\nfrom books, author, publishing_house\nwhere publishing_house.publishing_house_id = books.fk_publishing_house_id AND author.author_id = books.fk_author_id\norder by books.book_id;")
				if err != nil {
					log.Fatal("Did`not query:", err)
				}

				for rows.Next() {
					err1 := rows.Scan(&id, &book_name, &author_name, &publisher)
					if err1 != nil {
						log.Fatal("Did`not scan: ", err1)
					}
					log.Println(id, book_name.String, author_name.String, publisher.String)
				}
				if err != nil {
					fmt.Println("Can not:", err)
				}
				defer rows.Close()
			case "publishing_house":
				db, err := sql.Open("postgres", psqlconn)
				if err != nil {
					fmt.Println("Could not connect", err)
				}
				defer db.Close()
				var (
					id   int
					name sql.NullString
					city sql.NullString
				)
				rows, err := db.Query("select * from publishing_house")
				if err != nil {
					log.Fatal("Did`not query:", err)
				}

				for rows.Next() {
					err1 := rows.Scan(&id, &name, &city)
					if err1 != nil {
						log.Fatal("Did`not scan: ", err1)
					}
					log.Println(id, name.String, city.String)
				}
				if err != nil {
					fmt.Println("Can not:", err)
				}
				defer rows.Close()
			}
		case 6:
			variable = false
			fmt.Println("Bye bye")
		}
	}

}
