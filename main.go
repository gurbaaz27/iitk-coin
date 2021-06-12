package main

import (
	"github.com/gurbaaz27/iitk-coin/controllers"
	"github.com/gurbaaz27/iitk-coin/database"
)

func main() {
	database.InitialiseDB()

	controllers.HandleRequests()
}
