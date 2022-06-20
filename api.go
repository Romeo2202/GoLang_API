package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

//My database structure
type User struct {
	Id       int    `json:"id"`
	TagID    string `json:"tagid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Points   int    `json:"points"`
}

func main() {
	//Creates Simple HTTP Sites
	http.HandleFunc("/GetAllAccounts", GetAllAccounts)
	http.HandleFunc("/GetUserID", GetUserID)
	http.HandleFunc("/GetUsername", GetUsername)
	http.HandleFunc("/GetUserPoints", GetUserPoints)
	http.HandleFunc("/Updatepoints", UpdatePoints)

	//Starts API on port 10000
	log.Fatal(http.ListenAndServe(":10000", nil))

}

//Creates DB Connection and retrieves all account information
func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	var myUser User
	db, err := sql.Open("mysql", "root:Welkom01!@tcp(192.168.1.13:3306)/dba")

	//Simple error handling
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	//Data Selection
	results, err := db.Query("SELECT * FROM accounts")
	if err != nil {
		log.Fatal(err.Error())
	}
	//Converts data to JSON
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var data []User

	for results.Next() {
		results.Scan(&myUser.Id, &myUser.TagID, &myUser.Username, &myUser.Password, &myUser.Points)
		fmt.Println(myUser)
		data = append(data, myUser)
	}
	json.NewEncoder(w).Encode(data)
	fmt.Print(data)
}

//Creates DB Connection and retrieves userid
func GetUserID(w http.ResponseWriter, r *http.Request) {
	var myUser User
	db, err := sql.Open("mysql", "root:Welkom01!@tcp(192.168.1.13:3)/dba")

	//Simple error handling
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	//Data Selection
	results, err := db.Query("SELECT id FROM accounts")
	if err != nil {
		log.Fatal(err.Error())
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var data []User

	//Converts data to JSON
	for results.Next() {
		results.Scan(&myUser.Id)
		fmt.Println(myUser)
		data = append(data, myUser)
	}
	json.NewEncoder(w).Encode(data)
	fmt.Print(data)
}

//Creates DB Connection and retrieves username
func GetUsername(w http.ResponseWriter, r *http.Request) {
	var myUser User
	db, err := sql.Open("mysql", "root:Welkom01!@tcp(192.168.1.13:3306)/dba")

	//Simple error handling
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	//Data Selection
	results, err := db.Query("SELECT username FROM accounts")
	if err != nil {
		log.Fatal(err.Error())
	}
	w.WriteHeader(http.StatusCreated)
	var data []User

	//Converts data to JSON
	for results.Next() {
		results.Scan(&myUser.Username)
		fmt.Println(myUser)
		data = append(data, myUser)
	}
	json.NewEncoder(w).Encode(data)
	fmt.Print(data)
}

//Creates DB Connection and updates specific user
func GetUserPoints(w http.ResponseWriter, r *http.Request) {
	var myUser User
	db, err := sql.Open("mysql", "root:Welkom01!@tcp(192.168.1.13:3306)/dba")

	//Simple error handling
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	//Data Selection
	results, err := db.Query("SELECT points FROM accounts")
	if err != nil {
		log.Fatal(err.Error())
	}

	//Converts data to JSON
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	var data []User

	for results.Next() {
		results.Scan(&myUser.Points)
		fmt.Println(myUser)
		data = append(data, myUser)
	}
	json.NewEncoder(w).Encode(data)
	fmt.Print(data)
}

func UpdatePoints(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:Welkom01!@tcp(192.168.1.13:3306)/dba")

	//Simple error handling
	if err != nil {
		log.Fatal(err.Error())
	}

	defer db.Close()

	//listen to post request
	keys, ok := r.URL.Query()["id"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	key := keys[0]
	//Update points
	update, err := db.Query("UPDATE accounts SET points = points + 10 WHERE tagid = '" + key + "'")

	if err != nil {
		log.Fatal(err.Error())
	}
	defer update.Close()
}
