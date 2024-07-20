package main

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

type Item struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
}

var items []Item

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.NotFound(w, r)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	items = append(items, item)
	json.NewEncoder(w).Encode(item)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			var newItem Item
			_ = json.NewDecoder(r.Body).Decode(&newItem)
			newItem.ID = params["id"]
			items = append(items, newItem)
			json.NewEncoder(w).Encode(newItem)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Sample data
	items = append(items, Item{ID: "1", Name: "Item One", Price: 10.00})
	items = append(items, Item{ID: "2", Name: "Item Two", Price: 20.00})

	// Route handlers & endpoints
	r.HandleFunc("/items", getItems).Methods("GET")
	r.HandleFunc("/items/{id}", getItem).Methods("GET")
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	// Start server
	http.ListenAndServe(":8000", r)
}
