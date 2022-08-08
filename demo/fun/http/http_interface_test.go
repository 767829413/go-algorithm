package http

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

//func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	for item, price := range db {
//		fmt.Fprintf(w, "%s: %s\n", item, price)
//	}
//}

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			msg := fmt.Sprintf("no such page: %s\n", req.URL)
			http.Error(w, msg, http.StatusNotFound) // 404
			return
		}
		fmt.Fprintf(w, "%s\n", price)
	default:
		msg := fmt.Sprintf("no such page: %s\n", req.URL)
		http.Error(w, msg, http.StatusNotFound) // 404
	}
}

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound) // 404
		fmt.Fprintf(w, "no such item: %q\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}

func TestEg1(t *testing.T) {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func TestEg2(t *testing.T) {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}
