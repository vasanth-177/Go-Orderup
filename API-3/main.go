package main

import (
	"database/sql"
	"encoding/json"

	"encoding/base64"
	"fmt"
	"io/ioutil"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	// "log"
	"bytes"
	"net/http"
	"os"
	"strconv"
)

type Demo struct {
	Name         string `json:"name"`
	OrderDetails []struct {
		Item     string `json:"item"`
		HName    string `json:"h_name"`
		Quantity string `json:"quantity"`
	} `json:"orderDetails"`
}

type Dbdata struct {
	Email   string `json:"mail"`
	Contact string `json:"contact"`
	Address string `json:"address"`
}

// type Demo struct {
// 	Item     string `json:"item"`
// 	HName    string `json:"h_name"`
// 	Quantity string `json:"quantity"`
// }

var cart = map[string]map[string]uint64{}
var uname string
var email string
var contact string
var address string

func Gbill(w http.ResponseWriter, r *http.Request) {
	var temp = map[string]map[string]uint64{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	u := Demo{}
	json.Unmarshal([]byte(body), &u)
	uname = u.Name
	s := "["
	for i := 0; i < len(u.OrderDetails); i++ {
		s1 := ""
		if i == len(u.OrderDetails)-1 {
			s1 = `{"item":"` + u.OrderDetails[i].Item + `","h_name":"` + u.OrderDetails[i].HName + `","quantity":"` + u.OrderDetails[i].Quantity + `"}`
		} else {
			s1 = `{"item":"` + u.OrderDetails[i].Item + `","h_name":"` + u.OrderDetails[i].HName + `","quantity":"` + u.OrderDetails[i].Quantity + `"},`
		}

		temp[u.OrderDetails[i].Item] = map[string]uint64{}
		number, _ := strconv.ParseUint(u.OrderDetails[i].Quantity, 10, 64)
		temp[u.OrderDetails[i].Item][u.OrderDetails[i].HName] = number
		s = s + s1
	}
	s = s + "]"
	cart = temp
	extra_info(uname)
	generate_bill()
	sendMail()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	uuid := time.Now().Format("2006-01-02 15:04:05 Monday")

	se := base64.StdEncoding.EncodeToString([]byte(uuid))

	insert, err := db.Query("INSERT INTO orders value(?,?,?,?)", se, uname, time.Now().Format("2006-01-02 15:04:05 Monday"), s)
	// fmt.Println(reg1.Mail)
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	// fmt.Fprintf(w, "Bill generated successfully...")

	filename := "/Users/vasanth/golang-demo/API-3/pdfs/" + uname + ".pdf"
	streamPDFbytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	b := bytes.NewBuffer(streamPDFbytes)
	w.Header().Set("Content-type", "application/pdf")
	if _, err := b.WriteTo(w); err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	w.Write([]byte("PDF Generated"))
	// })

	// err := http.ListenAndServe(":4111", nil)
	// if err != nil {
	//     log.Fatal("ListenAndServe: ", err)
	//     fmt.Println(err)
	// }

}

// func Gbill(w http.ResponseWriter, r *http.Request) {
// 	var temp = map[string]map[string]uint64{}
// 	uname = r.URL.Query().Get("name")
// 	oD := r.URL.Query().Get("orderDetails")
// 	if len(uname) == 0 || len(oD) == 0 {
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
// 		fmt.Fprintf(w, "wrong request....")
// 	} else {
// 		// oD := `[{"item":"idly","h_name":"hotel1","quantity":"2"},{"item":"masala vada","h_name":"hotel1","quantity":"1"}]`
// 		// fmt.Println(oD)
// 		u := []Demo{}
// 		err := json.Unmarshal([]byte(oD), &u)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// fmt.Println(u)
// 		for i := 0; i < len(u); i++ {
// 			temp[u[i].Item] = map[string]uint64{}
// 			number, _ := strconv.ParseUint(u[i].Quantity, 10, 64)
// 			temp[u[i].Item][u[i].HName] = number
// 		}
// 		cart = temp
// 		generate_bill()
// 		// sendMail()
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT")
// 		fmt.Fprintf(w, "Bill generated successfully...")
// 	}
// }

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/generateBill", Gbill).Methods("POST")
	//myRouter.HandleFunc("/generateBill", Gbill).Methods("GET")
	// log.Fatal(http.ListenAndServe(":8083", myRouter))
	http.Handle("/", myRouter) // enable the router
	// Start the server.
	port := ":8083"
	fmt.Println("\nListening on port " + port)
	http.ListenAndServe(port, myRouter) // mux.Router now in play
}

func extra_info(uname string) {
	var dbdata Dbdata
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	result, err := db.Query(`select mail,contact,address from user where name = ?`, uname)
	if err != nil {
		panic(err.Error())
	}
	for result.Next() {
		err = result.Scan(&dbdata.Email, &dbdata.Contact, &dbdata.Address)
		if err != nil {
			panic(err.Error())
		}
	}
	email = dbdata.Email
	contact = dbdata.Contact
	address = dbdata.Address
}

func main() {
	handleRequests()
}
