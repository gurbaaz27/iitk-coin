package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gurbaaz27/iitk-coin/database"
	"github.com/gurbaaz27/iitk-coin/models"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to Home Page of IITK Coin. There are 3 useful endpoints are:- /signup, /login, /secretpage")
}

func secretPage(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Congrats! The fact that you have reached here is proof that you are successfully logged in!"))
}

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
		// Get(MyDB)
		w.Write([]byte("Welcome to Signup Page!\nSend a POST request to signup in iitkcoin.\n"))

	case "POST":
		var newUser models.User
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&newUser)
		checkError(err)
		log.Println(newUser)
		database.AddUser(database.MyDB, newUser)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newUser)

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}

func logIn(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/login" {
		http.NotFound(w, r)
		return
	}

	// db, _ := sql.Open("sqlite3", "./database.db")
	// defer db.Close()

	fmt.Println(r.URL.Path)
	switch r.Method {
	case "GET":
		//Get(db)
		w.Write([]byte("Welcome to Login Page!\nSend a POST request to login into iitkcoin.\n"))

	case "POST":

		var loginRequest models.LoginRequest
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&loginRequest)
		checkError(err)
		log.Println(loginRequest)
		if UserValid(db, loginRequest) {

			// Declare the expiration time of the token
			// here, we have kept it as 5 minutes
			expirationTime := time.Now().Add(10 * time.Minute)
			// Create the JWT claims, which includes the username and expiry time
			claims := &CustomClaims{
				Rollno: loginRequest.Rollno,
				StandardClaims: jwt.StandardClaims{
					// In JWT, the expiry time is expressed as unix milliseconds
					ExpiresAt: expirationTime.Unix(),
				},
			}

			// Declare the token with the algorithm used for signing, and the claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			// Create the JWT string
			tokenString, err := token.SignedString(jwtKey)
			CheckError(err)
			log.Println("TOKEN:", tokenString)
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})
			w.Write([]byte("U son of a bitch. I am in \n"))
		} else {
			w.Write([]byte("F....Invalid user credentials.\n"))
		}

	default:
		w.WriteHeader(http.StatusNotImplemented)
		w.Write([]byte(http.StatusText(http.StatusNotImplemented)))
	}

}

func HandleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/login", logIn)
	http.HandleFunc("/secretpage", secretPage)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
