package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	rollno int
	name   string
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func insertUserData(db *sql.DB, user User) {
	statement, err := db.Prepare("INSERT INTO User (rollno, name) VALUES (?, ?)")
	checkErr(err)
	statement.Exec(user.rollno, user.name)

	fmt.Printf("New user details : rollno = %d, name = %s added in database iitkcoin-190349.db\n", user.rollno, user.name)
}

func main() {
	db, err := sql.Open("sqlite3", "./iitkcoin-190349.db")
	checkErr(err)

	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS User (id INTEGER PRIMARY KEY, rollno INTEGER, name TEXT)")
	checkErr(err)
	statement.Exec()

	dummyuser1 := User{190349, "Harry Potter"}
	dummyuser2 := User{190500, "James Bond"}
	dummyuser3 := User{190064, "Wonder Woman"}

	insertUserData(db, dummyuser1)
	insertUserData(db, dummyuser2)
	insertUserData(db, dummyuser3)

	db.Close()
}
