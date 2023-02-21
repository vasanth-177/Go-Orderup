package main

import (
	// "database/sql"
	// "encoding/json"
	"database/sql"
	"encoding/json"
	"fmt"

	// "io/ioutil"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	// "log"
	"net/http"
)

type Response struct {
	Option string `json:"option"`
}

type Product struct {
	Item  string  `json:"item"`
	Price float64 `json:"price"`
}

func helper(s string) []Product {
	var products []Product
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	s1 := "SELECT * FROM " + s
	result, err := db.Query(s1)
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var product Product
		err = result.Scan(&product.Item, &product.Price)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	return products
}

func get_hnames() []string {
	hotelSlice := []string{}
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	result, err := db.Query("SELECT h_name FROM info ")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var product Product
		err = result.Scan(&product.Item)
		if err != nil {
			panic(err.Error())
		}
		hotelSlice = append(hotelSlice, product.Item)
	}
	return hotelSlice
}

// func GetInp(w http.ResponseWriter, r *http.Request) {
// 	opt := r.URL.Query().Get("type")
// 	// opt := mux.Vars(r)["type"]
// 	// fmt.Println(opt)
// 	var products []Product
// 	if len(opt) > 0 {
// 		hotelSlice := get_hnames()
// 		for _, v := range hotelSlice {
// 			if v == opt {
// 				products = helper(opt)
// 				break
// 			}
// 		}
// 	} else {

// 		reqBody, _ := ioutil.ReadAll(r.Body)
// 		var response Response
// 		json.Unmarshal(reqBody, &response)
// 		products = helper(response.Option)
// 	}

// 	// fmt.Fprintf(w, "%+v", string(reqBody))

// 	// fmt.Println(string(reqBody))

// 	w.Header().Set("Content-Type", "application/json")
// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500")
// 	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
// 	json.NewEncoder(w).Encode(products)

// 	fmt.Println("Hotels details sent successfully...")
// }

func GetItem(w http.ResponseWriter, r *http.Request) {
	opt := r.URL.Query().Get("type")
	var products []Product
	if len(opt) > 0 {
		hotelSlice := get_hnames()
		for _, v := range hotelSlice {
			if v == opt {
				products = helper(opt)
				break
			}
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	json.NewEncoder(w).Encode(products)

	fmt.Println("Hotels details sent successfully...")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/postOption", GetInp).Methods("POST")
	myRouter.HandleFunc("/getItem", GetItem).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8082", myRouter))
	http.Handle("/", myRouter) // enable the router
	// Start the server.
	port := ":8082"
	fmt.Println("\nListening on port " + port)
	http.ListenAndServe(port, myRouter) // mux.Router now in play
}

func main() {
	handleRequests()
}

// code := r.URL.Query().Get("type")
// code := mux.Vars(r)["type"]
