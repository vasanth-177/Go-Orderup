package main

import (
	// "io"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	// "strings"
	"testing"
)

func TestGetItem(t *testing.T) {
	tt := []struct {
		url      string
		dbOption string
	}{
		{url: `http://localhost:8082/getItem?type=hotel1`, dbOption: "hotel1"},
		{url: `http://localhost:8082/getItem?type=hotel2`, dbOption: "hotel2"},
		{url: `http://localhost:8082/getItem?type=hotel3`, dbOption: "hotel3"},
		{url: `http://localhost:8082/getItem?type=hotel4`, dbOption: "hotel4"},
		{url: `http://localhost:8082/getItem?type=hotel5`, dbOption: "hotel5"},
		{url: `http://localhost:8082/getItem`, dbOption: ""},
	}
	for _, tc := range tt {
		var products []Product
		hotelSlice := get_hnames()
		for _, v := range hotelSlice {
			if v == tc.dbOption {
				products = helper(tc.dbOption)
				break
			}
		}
		req, err := http.NewRequest("GET", tc.url, nil)
		if err != nil {
			t.Fatalf("could not create request:%v", err)
		}
		rec := httptest.NewRecorder()
		GetItem(rec, req)
		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != 200 {
			t.Errorf("expected status OK,got %v", res.Status)
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("coul not read response:%v", err)
		}
		var u []Product
		json.Unmarshal([]byte(b), &u)
		// fmt.Println(options, u)
		flag := reflect.DeepEqual(products, u)
		if !flag {
			t.Error("Wrong output")
		}
	}
}
