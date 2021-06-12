package database

import (
	"database/sql"
	"log"

	"github.com/gurbaaz27/iitk-coin/models"
	_ "github.com/mattn/go-sqlite3"
	_ "golang.org/x/crypto/bcrypt"
)

var MyDB *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitialiseDB() {
	MyDB, err := sql.Open("sqlite3", "./iitkcoin-190349.MyDB")
	checkErr(err)
	statement, err := MyDB.Prepare("CREATE TABLE IF NOT EXISTS User (id INTEGER PRIMARY KEY, rollno INTEGER, name TEXT, password TEXT)")
	checkErr(err)
	statement.Exec()
	defer MyDB.Close()
}

func AddUser(user models.User) {

	if UserExists(user) {
		statement, err := MyDB.Prepare("INSERT INTO User (rollno, name, password) VALUES (?, ?, ?)")
		checkErr(err)
		statement.Exec(user.Rollno, user.Name, database.HashPwd(user.Password))

		log.Println("New user details : rollno = %d, name = %s added in database iitkcoin-190349.MyDB\n", user.Rollno, user.Name)
	} else {
		log.Println("User with same roll no. already exists!")
	}
}

func UserValid(user models.LoginRequest) bool {
	// result := MyDB.Where

	return true
}

func UserExists(user models.User) bool {
	err := MyDB.QueryRow("`SELECT rollno FROM User WHERE rollno = ?`").Scan(&user.Rollno)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

// func
