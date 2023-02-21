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

func TestGbill(t *testing.T) {
	payload := strings.NewReader(`{"name":"vasanth","orderDetails":[{"item":"idly","h_name":"hotel1","quantity":"1"},{"item":"masala vada","h_name":"hotel1","quantity":"1"}]}`)
	req, err := http.NewRequest("POST", "http://localhost:8083/generateBill", payload)
	if err != nil {
		t.Fatalf("could not create request:%v", err)
	}
	rec := httptest.NewRecorder()
	Gbill(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Errorf("expected status OK,got %v", res.Status)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("coul not read response:%v", err)
	}
	if string(b) != "Bill generated successfully..." {
		t.Error("Wrong output")
	}
}

// func TestGbill(t *testing.T) {
// 	tt := []struct {
// 		url      string
// 		expected string
// 	}{
// 		{url: `http://localhost:8083/generateBill?name=vasanth&orderDetails=[{"item":"idly","h_name":"hotel1","quantity":"2"},{"item":"masala vada","h_name":"hotel1","quantity":"1"}]`, expected: "Bill generated successfully..."},
// 		{url: `http://localhost:8083/generateBill`, expected: "wrong request...."},
// 	}
// 	for _, tc := range tt {
// 		req, err := http.NewRequest("GET", tc.url, nil)
// 		if err != nil {
// 			t.Fatalf("could not create request:%v", err)
// 		}
// 		rec := httptest.NewRecorder()
// 		Gbill(rec, req)
// 		res := rec.Result()
// 		defer res.Body.Close()
// 		if res.StatusCode != 200 {
// 			t.Errorf("expected status OK,got %v", res.Status)
// 		}
// 		b, err := ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			t.Fatalf("coul not read response:%v", err)
// 		}
// 		if string(b) != tc.expected {
// 			t.Error("Wrong output")
// 		}
// 	}
// }
