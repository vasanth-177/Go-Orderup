package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	// "io/ioutil"
	// "log"
	"net/http"
)

type Resp struct {
	Status string `json:"status"`
}
type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Reg struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Mail     string `json:"email"`
	Contact  string `json:"contact"`
	Address  string `json:"address"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var flag bool = false
	var resp Resp
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user1 User
	json.Unmarshal(reqBody, &user1)
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Query("SELECT name FROM user")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var user User
		err = result.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		if user.Name == user1.Name {

			pwd, err := db.Query("SELECT password FROM user WHERE name = ?", user.Name)
			if err != nil {
				panic(err.Error())
			}
			for pwd.Next() {
				err = pwd.Scan(&user.Password)
				if err != nil {
					panic(err.Error())
				}
				if user.Password == user1.Password {
					flag = true
					break
				} else {
					break
				}
			}
		}
	}
	if flag {
		resp.Status = "true"
		json.NewEncoder(w).Encode(resp)
	} else {
		resp.Status = "false"
		json.NewEncoder(w).Encode(resp)
	}
	fmt.Println("Info sent")
}

// func Login(w http.ResponseWriter, r *http.Request) {
// 	var flag bool = false
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
// 	// reqBody, _ := ioutil.ReadAll(r.Body)
// 	// var user1 User
// 	// json.Unmarshal(reqBody, &user1)
// 	name := r.URL.Query().Get("name")
// 	password := r.URL.Query().Get("password")
// 	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	defer db.Close()
// 	result, err := db.Query("SELECT name FROM user")
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	for result.Next() {
// 		var user User
// 		err = result.Scan(&user.Name)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		if user.Name == name {
// 			pwd, err := db.Query("SELECT password FROM user WHERE name = ?", user.Name)
// 			if err != nil {
// 				panic(err.Error())
// 			}
// 			for pwd.Next() {
// 				err = pwd.Scan(&user.Password)
// 				if err != nil {
// 					panic(err.Error())
// 				}
// 				if user.Password == password {
// 					flag = true
// 					break
// 				} else {
// 					break
// 				}
// 			}
// 		}
// 	}
// 	if flag {
// 		fmt.Fprintf(w, "1")
// 	} else {
// 		fmt.Fprintf(w, "0")
// 	}
// }

func Register(w http.ResponseWriter, r *http.Request) {
	var flag bool = true
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var reg1 Reg
	var resp Resp
	json.Unmarshal(reqBody, &reg1)

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Query("SELECT name FROM user")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var user User
		err = result.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		if user.Name == reg1.Name {
			flag = false
			break
		}
	}
	if flag {

		insert, err := db.Query("INSERT INTO user VALUES(?,?,?,?,?)", reg1.Name, reg1.Password, reg1.Mail, reg1.Address, reg1.Contact)
		// fmt.Println(reg1.Mail)
		if err != nil {
			panic(err.Error())
		}

		defer insert.Close()
		resp.Status = "true"
		json.NewEncoder(w).Encode(resp)
	} else {
		resp.Status = "false"
		json.NewEncoder(w).Encode(resp)
	}
}

func adminlogin(w http.ResponseWriter, r *http.Request) {
	var flag bool = false
	var resp Resp
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user1 User
	json.Unmarshal(reqBody, &user1)
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Query("SELECT name FROM adminInfo")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var user User
		err = result.Scan(&user.Name)
		if err != nil {
			panic(err.Error())
		}
		if user.Name == user1.Name {

			pwd, err := db.Query("SELECT password FROM adminInfo WHERE name = ?", user.Name)
			if err != nil {
				panic(err.Error())
			}
			for pwd.Next() {
				err = pwd.Scan(&user.Password)
				if err != nil {
					panic(err.Error())
				}
				if user.Password == user1.Password {
					flag = true
					break
				} else {
					break
				}
			}
		}
	}
	if flag {
		resp.Status = "true"
		json.NewEncoder(w).Encode(resp)
	} else {
		resp.Status = "false"
		json.NewEncoder(w).Encode(resp)
	}
	fmt.Println("Info sent")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/login", Login).Methods("POST")
	myRouter.HandleFunc("/register", Register).Methods("POST")
	myRouter.HandleFunc("/adminlogin", adminlogin).Methods("POST")
	// log.Fatal(http.ListenAndServe(":8080", myRouter))
	http.Handle("/", myRouter) // enable the router

	// Start the server.
	port := ":8080"
	fmt.Println("\nListening on port " + port)
	http.ListenAndServe(port, myRouter) // mux.Router now in play
}

func main() {
	handleRequests()
}

// code := r.URL.Query().Get("type")
// code := mux.Vars(r)["type"]
