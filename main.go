package main

import (
	"github.com/gurbaaz27/iitk-coin/controllers"
	"github.com/gurbaaz27/iitk-coin/database"
	_ "github.com/mattn/go-sqlite3"
	_ "golang.org/x/crypto/bcrypt"
)

func main() {
	database.InitialiseDB()

	controllers.HandleRequests()

	// dummyuser1 := User{190349, "Harry Potter"}
	// dummyuser2 := User{190500, "James Bond"}
	// dummyuser3 := User{190064, "Wonder Woman"}

	// insertUserData(db, dummyuser1)
	// insertUserData(db, dummyuser2)
	// insertUserData(db, dummyuser3)
}
