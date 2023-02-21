package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	// "io/ioutil"
	// "log"
	"net/http"
)

type Disp struct {
	Name string `json:"h_name"`
	Type int    `json:"type"`
}

type Option struct {
	Hotelname string `"json:h_name"`
	Type      string `"json:type"`
}

func helper(i int, s string) []Option {
	var options []Option
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	result, err := db.Query("SELECT h_name,type FROM info WHERE type "+s, i)
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var disp Disp
		var option Option
		err = result.Scan(&disp.Name, &disp.Type)
		if err != nil {
			panic(err.Error())
		}
		option.Hotelname = disp.Name
		if disp.Type == 0 {
			option.Type = "Veg hotel"
		} else if disp.Type == 1 {
			option.Type = "Non-veg hotel"
		} else {
			option.Type = "Veg/Non-veg hotel"
		}
		options = append(options, option)
	}
	return options
}

func ShowVeg(w http.ResponseWriter, r *http.Request) {
	options := helper(0, "= ?")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	json.NewEncoder(w).Encode(options)

	fmt.Println("Info sent successfully....")
}

func ShowNonVeg(w http.ResponseWriter, r *http.Request) {
	options := helper(1, "= ?")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	json.NewEncoder(w).Encode(options)

	fmt.Println("Info sent successfully....")
}

func ShowBoth(w http.ResponseWriter, r *http.Request) {
	options := helper(2, "= ?")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	json.NewEncoder(w).Encode(options)

	fmt.Println("Info sent successfully....")
}

func ShowAll(w http.ResponseWriter, r *http.Request) {
	opt := r.URL.Query().Get("type")
	var options []Option
	if len(opt) > 0 {
		if opt == "veg" {
			options = helper(0, "= ?")
		} else if opt == "non-veg" {
			options = helper(1, "= ?")
		} else if opt == "veg/non-veg" {
			options = helper(2, "= ?")
		}
	} else {
		options = helper(5, "!= ?")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	json.NewEncoder(w).Encode(options)

	fmt.Println("Info sent successfully....")

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/getHotels", ShowAll).Methods("GET")
	myRouter.HandleFunc("/getVegHotels", ShowVeg).Methods("GET")
	myRouter.HandleFunc("/getNonVegHotels", ShowNonVeg).Methods("GET")
	myRouter.HandleFunc("/getBothHotels", ShowBoth).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8081", myRouter))
	http.Handle("/", myRouter) // enable the router

	// Start the server.
	port := ":8081"
	fmt.Println("\nListening on port " + port)
	http.ListenAndServe(port, myRouter) // mux.Router now in play
}

func main() {
	handleRequests()
}

// code := r.URL.Query().Get("type")
// code := mux.Vars(r)["type"]
