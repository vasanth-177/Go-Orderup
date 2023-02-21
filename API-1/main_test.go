package main

import (
	// "io"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestShowNonVeg(t *testing.T) {
	options := helper(1, "= ?")
	req, err := http.NewRequest("GET", "http://localhost:8081/getBothHotels", nil)
	if err != nil {
		t.Fatalf("could not create request:%v", err)
	}
	rec := httptest.NewRecorder()
	ShowNonVeg(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Errorf("expected status OK,got %v", res.Status)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("coul not read response:%v", err)
	}
	var u []Option
	json.Unmarshal([]byte(b), &u)
	// fmt.Println(options, u)
	flag := reflect.DeepEqual(options, u)
	if !flag {
		t.Error("Wrong output")
	}

}

func TestShowVeg(t *testing.T) {
	options := helper(0, "= ?")
	req, err := http.NewRequest("GET", "http://localhost:8081/getVegHotels", nil)
	if err != nil {
		t.Fatalf("could not create request:%v", err)
	}
	rec := httptest.NewRecorder()
	ShowVeg(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Errorf("expected status OK,got %v", res.Status)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("coul not read response:%v", err)
	}
	var u []Option
	json.Unmarshal([]byte(b), &u)
	// fmt.Println(options, u)
	flag := reflect.DeepEqual(options, u)
	if !flag {
		t.Error("Wrong output")
	}

}

func TestShowBoth(t *testing.T) {
	options := helper(2, "= ?")
	req, err := http.NewRequest("GET", "http://localhost:8081/getBothHotels", nil)
	if err != nil {
		t.Fatalf("could not create request:%v", err)
	}
	rec := httptest.NewRecorder()
	ShowBoth(rec, req)
	res := rec.Result()
	defer res.Body.Close()
	if res.StatusCode != 200 {
		t.Errorf("expected status OK,got %v", res.Status)
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("coul not read response:%v", err)
	}
	var u []Option
	json.Unmarshal([]byte(b), &u)
	// fmt.Println(options, u)
	flag := reflect.DeepEqual(options, u)
	if !flag {
		t.Error("Wrong output")
	}

}

func TestShowAll(t *testing.T) {
	tt := []struct {
		url string
		id  int
		hs  string
	}{
		{url: `http://localhost:8081/getHotels`, id: 5, hs: "!= ?"},
		{url: `http://localhost:8081/getHotels?type=veg`, id: 0, hs: "= ?"},
		{url: `http://localhost:8081/getHotels?type=non-veg`, id: 1, hs: "= ?"},
		{url: `http://localhost:8081/getHotels?type=veg/non-veg`, id: 2, hs: "= ?"},
	}
	for _, tc := range tt {
		options := helper(tc.id, tc.hs)
		req, err := http.NewRequest("GET", tc.url, nil)
		if err != nil {
			t.Fatalf("could not create request:%v", err)
		}
		rec := httptest.NewRecorder()
		ShowAll(rec, req)
		res := rec.Result()
		defer res.Body.Close()
		if res.StatusCode != 200 {
			t.Errorf("expected status OK,got %v", res.Status)
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("coul not read response:%v", err)
		}
		var u []Option
		json.Unmarshal([]byte(b), &u)
		// fmt.Println(options, u)
		flag := reflect.DeepEqual(options, u)
		if !flag {
			t.Error("Wrong output")
		}
	}

}
