package main

import (
	// "io"
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogin(t *testing.T) {
	tt := []struct {
		name     string
		password string
		res      string
	}{
		{name: "vasanth", password: "12345", res: "1"},
		{name: "vasanth ", password: "1", res: "0"},
		{name: "kathir", password: "12345", res: "0"},
	}
	for _, tc := range tt {
		payload := strings.NewReader(`{
		"name":"` + tc.name + `",
		"password":"` + tc.password + `"
	}`)
		req, err := http.NewRequest("POST", "http://localhost:8080/login", payload)
		if err != nil {
			t.Fatalf("could not create request:%v", err)
		}
		rec := httptest.NewRecorder()
		Login(rec, req)
		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != 200 {
			t.Errorf("expected status OK,got %v", res.Status)
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("coul not read response:%v", err)
		}
		// fmt.Println(string(b))
		if string(b) != tc.res {
			t.Error("Wrong output")
		}
	}
}

// func TestLogin(t *testing.T) {
// 	tt := []struct {
// 		url string
// 		res string
// 	}{
// 		{url: "http://localhost:8080/login?name=vasanth&password=12345", res: "1"},
// 		{url: "http://localhost:8080/login?name=vasanth&password=1", res: "0"},
// 		{url: "http://localhost:8080/login?name=kathir&password=12345", res: "0"},
// 	}
// 	for _, tc := range tt {
// 		req, err := http.NewRequest("GET", tc.url, nil)
// 		if err != nil {
// 			t.Fatalf("could not create request:%v", err)
// 		}
// 		rec := httptest.NewRecorder()
// 		Login(rec, req)
// 		res := rec.Result()
// 		defer res.Body.Close()
// 		if res.StatusCode != 200 {
// 			t.Errorf("expected status OK,got %v", res.Status)
// 		}
// 		b, err := ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			t.Fatalf("coul not read response:%v", err)
// 		}
// 		// fmt.Println(string(b))
// 		if string(b) != tc.res {
// 			t.Error("Wrong output")
// 		}
// 	}
// }

func TestRegister(t *testing.T) {
	tt := []struct {
		name     string
		password string
		email    string
		address  string
		contact  string
		res      string
	}{
		{name: "vasanth", password: "12345", email: "vasanthmit17@gmail.com", address: "coimbatore", contact: "9360688299", res: "0"},
		{name: "kathir", password: "12345", email: "vasanthmit17@gmail.com", address: "coimbatore", contact: "9360688299", res: "1"},
	}
	for _, tc := range tt {
		payload := strings.NewReader(`{
			"name":"` + tc.name + `",
			"password":"` + tc.password + `",
			"email":"` + tc.email + `",
			"address":"` + tc.address + `",
			"contact":"` + tc.contact + `"
		}`)

		req, err := http.NewRequest("POST", "http://localhost:8080/register", payload)
		if err != nil {
			t.Fatalf("could not create request:%v", err)
		}
		rec := httptest.NewRecorder()
		Register(rec, req)
		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != 200 {
			t.Errorf("expected status OK,got %v", res.Status)
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("coul not read response:%v", err)
		}
		// fmt.Println(string(b))
		if string(b) != tc.res {
			t.Error("Wrong output detected....")
		}
	}
}
