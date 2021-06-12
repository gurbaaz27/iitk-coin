package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func signUp(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/signup" {
		http.NotFound(w, r)
		return
	}

	// database.db, _ := sql.Open("sqlite3", "./database.database.db")
	// defer database.db.Close()

	fmt.Println(r.URL.Path)
	switch r.Method {
	case "GET":
		Get(database.db)
		w.Write([]byte("Received a Get request\n"))

	case "POST":

		var newUser SignupJSON
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newUser)
		CheckError(err)
		log.Println(newUser)
		Add(database.db, newUser)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newUser)
	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login", logIn)
	http.HandleFunc("/secretpage", secretPage)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
