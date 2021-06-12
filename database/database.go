package database

import (
	"database/sql"
	"log"

	"github.com/gurbaaz27/iitk-coin/models"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
)

var MyDB *sql.DB

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func InitialiseDB() {
	MyDB, err := sql.Open("sqlite3", "./iitkcoin-190349.db")
	checkErr(err)
	defer MyDB.Close()
	statement, err := MyDB.Prepare("CREATE TABLE IF NOT EXISTS User (id INTEGER PRIMARY KEY, rollno INTEGER, name TEXT, password TEXT)")
	checkErr(err)
	log.Println("Database opened and table created successfully!")
	statement.Exec()
}

func AddUser(user models.User) {
	MyDB, err := sql.Open("sqlite3", "./iitkcoin-190349.MyDB")
	checkErr(err)
	defer MyDB.Close()
	if UserExists(user) {
		statement, err := MyDB.Prepare("INSERT INTO User (rollno, name, password) VALUES (?, ?, ?)")
		checkErr(err)
		statement.Exec(user.Rollno, user.Name, HashPwd(user.Password))

		log.Printf("New user details : rollno = %d, name = %s added in database iitkcoin-190349.MyDB\n", user.Rollno, user.Name)
	} else {
		log.Println("User with same roll no. already exists!")
	}
}

func UserValid(user models.LoginRequest) bool {
	MyDB, err := sql.Open("sqlite3", "./iitkcoin-190349.MyDB")
	checkErr(err)
	defer MyDB.Close()
	rows, err := MyDB.Query("SELECT * from User")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var rollno int64
		var name string
		var password string
		rows.Scan(&rollno, &name, &password)
		if user.Rollno == rollno && CheckPasswords(password, user.Password) {
			return true
		}
	}

	return false
}

func UserExists(user models.User) bool {
	err := MyDB.QueryRow("SELECT rollno FROM User WHERE rollno = ?").Scan(&user.Rollno)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return false
	}
	return true
}

func HashPwd(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	return string(hashedPassword)
}

func CheckPasswords(hashedPwd string, pwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(pwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}
