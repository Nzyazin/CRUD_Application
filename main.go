package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strings"
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
	rows, err := db.Query("SELECT id from author order by id desc limit 1")
	for rows.Next() {
		err1 := rows.Scan(&author_id)
		if err1 != nil {
			log.Fatal(err1)
		}
	}
	stmt, err := db.Prepare("INSERT INTO author(id, author_name, birthday) VALUES($1, $2, $3)")
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
		fmt.Println("Input a number for procedure with database or for quit: \n1. Add a row in table. \n2. Output all existing table in database. \n3. Read a table. \n4. Update a field in table. \n5. Delete a field. \n6. Quit")
		fmt.Scanln(&operation)
		psqlconn := fmt.Sprintf("host= %s port = %d user = %s password =%s dbname = %s sslmode=disable", host, port, user, password, dbname)
		switch operation {
		case 1:
			fmt.Println("Input name table (if you don`t know, you can use 2nd option): ")
			in := bufio.NewReader(os.Stdin)
			name_table, err := in.ReadString('\n')
			if err != nil {
				log.Println("Didn`t read name table: ", err)
			}
			name_table = strings.Join(strings.Fields(name_table), "")
			switch name_table {
			case "author":
				fmt.Println("Input author name:")
				in := bufio.NewReader(os.Stdin)
				name, err := in.ReadString('\n')
				if err != nil {
					log.Println(err)
				}
				fmt.Println("And his birthday in format <YYYY-MM-DD>:")
				var birthday string
				fmt.Scanln(&birthday)

				create(name, birthday)
			case "books":
				db, err := sql.Open("postgres", psqlconn)
				if err != nil {
					fmt.Println("Could not connect", err)
				}
				defer db.Close()
				var book_id int
				rows, err := db.Query("SELECT id from books order by id desc limit 1")
				for rows.Next() {
					err1 := rows.Scan(&book_id)
					if err1 != nil {
						log.Fatal(err1)
					}
				}
				var (
					b_name   string
					fk_ph_id int
					fk_a_id  int
				)
				fmt.Println("Input name book: ")
				fmt.Scanln(&b_name)
				b_name = strings.Join(strings.Fields(b_name), "")
				fmt.Println("Input foreign key from publisher id: ")
				fmt.Scanln(&fk_ph_id)
				fmt.Println("Input foreign key from author id: ")
				fmt.Scanln(&fk_a_id)
				stmt, err := db.Prepare("INSERT INTO author(id, book_name, fk_publishing_house_id, fk_author_id) VALUES($1, $2, $3, $4)")
				if err != nil {
					log.Fatal("Here1: ", err)
				}
				res, err := stmt.Exec(book_id+1, b_name, fk_ph_id, fk_a_id)
				if err != nil {
					log.Fatal("Here2: ", err)
				}
				log.Println(res)
			case "publishing_house":
				db, err := sql.Open("postgres", psqlconn)
				if err != nil {
					fmt.Println("Could not connect", err)
				}
				defer db.Close()
				var pb_id int
				rows, err := db.Query("SELECT id from publishing_house order by id desc limit 1")
				for rows.Next() {
					err1 := rows.Scan(&pb_id)
					if err1 != nil {
						log.Fatal(err1)
					}
				}
				var (
					pb_name string
					pb_city string
				)
				fmt.Println("Input name publisher: ")
				fmt.Scanln(&pb_name)
				pb_name = strings.Join(strings.Fields(pb_name), "")
				fmt.Println("Input name city of publisher: ")
				fmt.Scanln(&pb_city)
				pb_city = strings.Join(strings.Fields(pb_city), "")

				stmt, err := db.Prepare("INSERT INTO publishing_house(id, ph_name, city) VALUES($1, $2, $3)")
				if err != nil {
					log.Fatal("Here1: ", err)
				}
				res, err := stmt.Exec(pb_id+1, pb_name, pb_city)
				if err != nil {
					log.Fatal("Here2: ", err)
				}
				log.Println(res)
			}

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
				rows, err := db.Query("select books.id, books.book_name, author.author_name as author, publishing_house.ph_name as publisher\nfrom books, author, publishing_house\nwhere publishing_house.id = books.fk_publishing_house_id AND author.id = books.fk_author_id\norder by books.id;")
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
		case 4:
			fmt.Println("Input name of table for update <author, books, publishing_house>: ")
			in := bufio.NewReader(os.Stdin)
			name_table, err := in.ReadString('\n')
			if err != nil {
				log.Println(err)
			}
			name_table = name_table[:len(name_table)-2]
			switch name_table {
			case "author":
				db, err := sql.Open("postgres", psqlconn)
				if err != nil {
					fmt.Println("Could not connect", err)
				}
				defer db.Close()
				var (
					column_name sql.NullString
				)
				rows, err := db.Query("select column_name\n    from INFORMATION_SCHEMA.COLUMNS \n    where table_name = 'author';")
				for rows.Next() {
					err11 := rows.Scan(&column_name)
					if err11 != nil {
						fmt.Println("Can`t scan <author>")
					}
					log.Println(column_name.String)
				}
				fmt.Println("Input name of column for update <author_name, birthday>: ")
				in1 := bufio.NewReader(os.Stdin)
				name_column, err1 := in1.ReadString('\n')
				if err != nil {
					log.Println(err1)
				}
				/*
					var old_var string
					fmt.Fscan(os.Stdin, &old_var)*/
				fmt.Println("Input new variable of column for update: ")
				in2 := bufio.NewReader(os.Stdin)
				new_varya, err2 := in2.ReadString('\n')
				new_varya = strings.Join(strings.Fields(new_varya), " ")
				if err != nil {
					log.Println(err2)
				}
				fmt.Println("Input old name of object for update: ")
				in3 := bufio.NewReader(os.Stdin)
				old_var, err3 := in3.ReadString('\n')
				old_var = strings.Join(strings.Fields(old_var), " ")
				if err3 != nil {
					log.Println(err3)
				}

				query := fmt.Sprintf("UPDATE %s SET %s = '%s' WHERE %s = '%s';", name_table, name_column, new_varya, name_column, old_var)
				fmt.Println(query)

				res, err := db.Exec(query)
				if err != nil {
					log.Fatal("Did`not exec:", err)
				}
				log.Println(res)
			case "books":
				db, err := sql.Open("postgres", psqlconn)
				if err != nil {
					fmt.Println("Could not connect", err)
				}
				defer db.Close()
				var (
					column_name sql.NullString
				)
				rows, err := db.Query("select column_name\n    from INFORMATION_SCHEMA.COLUMNS \n    where table_name = 'books';")
				for rows.Next() {
					err11 := rows.Scan(&column_name)
					if err11 != nil {
						fmt.Println("Can`t scan <books>")
					}
					log.Println(column_name.String)
				}
				fmt.Println("Input name of column for update: ")
				in1 := bufio.NewReader(os.Stdin)
				name_column, err1 := in1.ReadString('\n')
				if err != nil {
					log.Println(err1)
				}
				name_column = strings.Join(strings.Fields(name_column), " ")
				//name_column = name_column[:len(name_column)-2]
				switch name_column {
				case "book_name":
					fmt.Println("Input new variable of field for update:")
					in1 := bufio.NewReader(os.Stdin)
					new_varya, err1 := in1.ReadString('\n')
					if err != nil {
						log.Println(err1)
					}
					fmt.Println("Input old name of field for update:")
					in2 := bufio.NewReader(os.Stdin)
					old_var, err1 := in2.ReadString('\n')
					if err != nil {
						log.Println(err1)
					}
					query := fmt.Sprintf("UPDATE %s SET %s = '%s' WHERE %s = '%s';", name_table, name_column, new_varya, name_column, old_var)
					fmt.Println(query)
					res, err := db.Exec(query)
					if err != nil {
						log.Fatal("Did`not exec:", err)
					}
					log.Println(res)
				default:
					fmt.Println("Input new variable of field for update: ")
					var new_varya int
					fmt.Scanln(&new_varya)
					fmt.Println("Input old name of field for update: ")
					var old_var int
					fmt.Scanln(&old_var)
					query := fmt.Sprintf("UPDATE %s SET %s = %d WHERE %s = %d;", name_table, name_column, new_varya, name_column, old_var)
					fmt.Println(query)
					res, err := db.Exec(query)
					if err != nil {
						log.Fatal("Did`not exec:", err)
					}
					log.Println(res)
				}
			case "publishing_house":
				db, err := sql.Open("postgres", psqlconn)
				if err != nil {
					fmt.Println("Could not connect", err)
				}
				defer db.Close()
				var (
					column_name sql.NullString
				)
				rows, err := db.Query("select column_name\n    from INFORMATION_SCHEMA.COLUMNS \n    where table_name = 'publishing_house';")
				for rows.Next() {
					err11 := rows.Scan(&column_name)
					if err11 != nil {
						fmt.Println("Can`t scan <publishing_house>")
					}
					log.Println(column_name.String)
				}
				fmt.Println("Input name of column for update (exclude <id>): ")
				in1 := bufio.NewReader(os.Stdin)
				name_column, err1 := in1.ReadString('\n')
				if err != nil {
					log.Println(err1)
				}
				name_column = strings.Join(strings.Fields(name_column), " ")
				switch name_column {
				case "id":
					fmt.Println("Input new variable of field for update: ")
					var new_varya int
					fmt.Scanln(&new_varya)
					fmt.Println("Input old name of field for update: ")
					var old_var int
					fmt.Scanln(&old_var)
					query := fmt.Sprintf("UPDATE %s SET %s = %d WHERE %s = %d;", name_table, name_column, new_varya, name_column, old_var)
					fmt.Println(query)
					res, err := db.Exec(query)
					if err != nil {
						log.Fatal("Did`not exec:", err)
					}
					log.Println(res)
				default:
					fmt.Println("Input new variable of field for update:")
					in1 := bufio.NewReader(os.Stdin)
					new_varya, err1 := in1.ReadString('\n')
					if err != nil {
						log.Println(err1)
					}
					new_varya = strings.Join(strings.Fields(new_varya), "")
					fmt.Println("Input old name of field for update:")
					in2 := bufio.NewReader(os.Stdin)
					old_var, err1 := in2.ReadString('\n')
					if err != nil {
						log.Println(err1)
					}
					old_var = strings.Join(strings.Fields(old_var), "")
					query := fmt.Sprintf("UPDATE %s SET %s = '%s' WHERE %s = '%s';", name_table, name_column, new_varya, name_column, old_var)
					fmt.Println(query)
					res, err := db.Exec(query)
					if err != nil {
						log.Fatal("Did`not exec:", err)
					}
					log.Println(res)
				}
			}
		case 5:
			db, err := sql.Open("postgres", psqlconn)
			if err != nil {
				fmt.Println("Could not connect", err)
			}
			defer db.Close()
			fmt.Println("Input name of table for deleting row <author, books, publishing_house>: ")
			in := bufio.NewReader(os.Stdin)
			name_table, err := in.ReadString('\n')
			if err != nil {
				log.Println(err)
			}
			name_table = name_table[:len(name_table)-2]
			fmt.Println("Input id of row: ")
			var table_id int
			fmt.Scanln(&table_id)
			if err != nil {
				log.Println(err)
			}
			query := fmt.Sprintf("DELETE FROM %s WHERE id = %d;", name_table, table_id)
			fmt.Println(query)
			res, err := db.Exec(query)
			if err != nil {
				log.Println(err)
			}
			fmt.Println(res)
		case 6:
			variable = false
			fmt.Println("Bye bye")
		}
	}

}
