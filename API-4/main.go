package main

import (
	// "database/sql"
	// "encoding/json"
	"database/sql"
	"encoding/json"
	"fmt"

	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/rs/cors"

	// "log"
	"net/http"
)

type Detail struct {
	Orderdetails string `json:"orderdetails"`
}

type Order struct {
	Date string `json:"date"`
	Id   string `json:"id"`
}

type Review struct {
	Name     string `json:"name"`
	Feedback string `json:"feedback"`
}
type Resp struct {
	Status string `json:"status"`
}

type Addhotel struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Additem struct {
	Name  string `json:"name"`
	Item  string `json:"item"`
	Price string `json:"price"`
}

func helper() []Review {
	var reviews []Review
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	result, err := db.Query("SELECT name,feedback FROM review")
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		var review Review
		err = result.Scan(&review.Name, &review.Feedback)
		if err != nil {
			panic(err.Error())
		}
		reviews = append(reviews, review)
	}
	return reviews
}

func SaveFeedback(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var review Review
	json.Unmarshal(reqBody, &review)

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO review value(?,?)", review.Name, review.Feedback)
	// fmt.Println(reg1.Mail)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	var resp Resp
	resp.Status = "true"
	json.NewEncoder(w).Encode(resp)

}

func GetFeedback(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	json.NewEncoder(w).Encode(helper())

}

func AddHotel(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var addhotel Addhotel
	json.Unmarshal(reqBody, &addhotel)

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if addhotel.Type == "Veg" {
		addhotel.Type = "0"
	} else if addhotel.Type == "Non-veg" {
		addhotel.Type = "1"
	} else {
		addhotel.Type = "2"
	}
	insert, err := db.Query("INSERT INTO info value(?,?)", addhotel.Name, addhotel.Type)
	// fmt.Println(reg1.Mail)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	cre, err := db.Exec("CREATE TABLE IF NOT EXISTS " + addhotel.Name + "(item varchar(255),price varchar(255))")
	if err != nil {
		panic(err)
	}
	fmt.Println(cre)
	var resp Resp
	resp.Status = "true"
	json.NewEncoder(w).Encode(resp)

}

func AddItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var additem Additem
	json.Unmarshal(reqBody, &additem)

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	que := "INSERT INTO " + additem.Name + " value(?,?)"
	insert, err := db.Query(que, additem.Item, additem.Price)
	// fmt.Println(reg1.Mail)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	var resp Resp
	resp.Status = "true"
	json.NewEncoder(w).Encode(resp)

}
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	// w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization")

	// if r.Method == "OPTIONS" {
	// 	w.WriteHeader(http.StatusNoContent)
	// 	return
	// }

	name := r.URL.Query().Get("name")
	item := r.URL.Query().Get("item")
	if len(name) > 0 && len(item) > 0 {
		s := "DELETE FROM " + name + " WHERE item = ?"
		delete, err := db.Query(s, item)
		if err != nil {
			panic(err)
		}
		defer delete.Close()
	} else {
		s := "DROP TABLE " + name
		_, err := db.Exec(s)
		if err != nil {
			panic(err)
		}
		s = "DELETE FROM info WHERE h_name = ?"
		delete, err := db.Query(s, name)
		if err != nil {
			panic(err)
		}
		defer delete.Close()
	}
	var resp Resp
	resp.Status = "true"
	json.NewEncoder(w).Encode(resp)
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	var orders []Order
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
	name := r.URL.Query().Get("name")
	id := r.URL.Query().Get("id")
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	if len(name) > 0 {
		result, err := db.Query("SELECT id,date FROM orders WHERE name = ?", name)
		if err != nil {
			panic(err.Error())
		}
		for result.Next() {
			var order Order
			err = result.Scan(&order.Id, &order.Date)
			if err != nil {
				panic(err.Error())
			}
			orders = append(orders, order)
		}

		json.NewEncoder(w).Encode(orders)
	}
	if len(id) > 0 {
		result, err := db.Query("SELECT orderdetails FROM orders WHERE id = ?", id)
		if err != nil {
			panic(err.Error())
		}
		var detail Detail
		for result.Next() {

			err = result.Scan(&detail.Orderdetails)
			if err != nil {
				panic(err.Error())
			}
		}

		json.NewEncoder(w).Encode(detail.Orderdetails)
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/saveFeedback", SaveFeedback).Methods("POST")
	myRouter.HandleFunc("/getFeedback", GetFeedback).Methods("GET")
	myRouter.HandleFunc("/addHotel", AddHotel).Methods("POST")
	myRouter.HandleFunc("/addItem", AddItem).Methods("POST")
	myRouter.HandleFunc("/deleteItem", DeleteItem).Methods("GET", "DELETE")

	myRouter.HandleFunc("/getOrders", GetOrders).Methods("GET")

	http.Handle("/", myRouter) // enable the router
	// Start the server.
	port := ":8084"
	fmt.Println("\nListening on port " + port)
	http.ListenAndServe(port, myRouter) // mux.Router now in play
}

func main() {
	handleRequests()
}
