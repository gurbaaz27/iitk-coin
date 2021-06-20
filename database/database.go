package database

import (
	"context"
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
	statement, err := MyDB.Prepare("CREATE TABLE IF NOT EXISTS User (rollno INTEGER, name TEXT, password TEXT, coins INTEGER)")
	checkErr(err)
	log.Println("Database opened and table created (if not existed) successfully!")
	statement.Exec()
}

func AddUser(user models.User) {
	MyDB, err := sql.Open("sqlite3", "./iitkcoin-190349.db")
	checkErr(err)
	defer MyDB.Close()
	if !UserExists(user) {
		statement, err := MyDB.Prepare("INSERT INTO User (rollno, name, password, coins) VALUES (?, ?, ?, ?)")
		checkErr(err)
		statement.Exec(user.Rollno, user.Name, HashPwd(user.Password), 0)

		log.Printf("New user details : rollno = %d, name = %s added in database iitkcoin-190349.db\n", user.Rollno, user.Name)
	} else {
		log.Println("User with same roll no. already exists!")
	}
}

func UserValid(user models.LoginRequest) bool {
	MyDB, err := sql.Open("sqlite3", "./iitkcoin-190349.db")
	checkErr(err)
	defer MyDB.Close()
	rows, err := MyDB.Query("SELECT * from User")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var rollno int64
		var name string
		var password string
		var coins int64
		err = rows.Scan(&rollno, &name, &password, &coins)
		checkErr(err)
		if user.Rollno == rollno && CheckPasswords(password, user.Password) {
			return true
		}
	}

	return false
}

func UserExists(user models.User) bool {
	MyDB, err := sql.Open("sqlite3", "./iitkcoin-190349.db")
	checkErr(err)
	defer MyDB.Close()
	err = MyDB.QueryRow("SELECT rollno FROM User WHERE rollno = ?", user.Rollno).Scan(&user.Rollno)
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

func ReturnBalance(rollno int64) int64 {
	MyDB, err := sql.Open("sqlite3", "./iitkcoin-190349.db")
	checkErr(err)
	defer MyDB.Close()
	var coins int64
	err = MyDB.QueryRow("SELECT coins FROM User WHERE rollno = ?", rollno).Scan(&coins)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return -1
	}
	return coins
}

func UpdateBalance(user models.RewardPayload) bool {
	MyDB, err := sql.Open("sqlite3", "./iitkcoin-190349.db")
	checkErr(err)
	defer MyDB.Close()
	log.Printf("Updating the balance!")
	statement, err := MyDB.Prepare("UPDATE User SET coins = coins + ? WHERE rollno = ?")
	if err != nil {
		return false
	}
	statement.Exec(user.Coins, user.Rollno)
	return true
}

func TransferCoins(user models.TransferPayload) bool {
	MyDB, err := sql.Open("sqlite3", "./iitkcoin-190349.db")
	checkErr(err)
	defer MyDB.Close()

	ctx := context.Background()
	tx, err := MyDB.BeginTx(ctx, nil)
	checkErr(err)

	res, err := tx.ExecContext(ctx, "UPDATE User SET coins = coins - ? WHERE rollno=? AND coins - ? >= 0", user.Coins, user.SenderRollno, user.Coins)
	checkErr(err)
	rows_affected, err := res.RowsAffected()
	checkErr(err)

	if rows_affected != 1 {
		tx.Rollback()
		return false
	}

	res, err = tx.ExecContext(ctx, "UPDATE User SET coins = coins + ? WHERE rollno=?", user.Coins, user.ReceiverRollno, user.Coins)
	checkErr(err)
	rows_affected, err = res.RowsAffected()
	checkErr(err)

	if rows_affected != 1 {
		tx.Rollback()
		return false
	}

	err = tx.Commit()
	checkErr(err)

	return true
}
