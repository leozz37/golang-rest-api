package main

import (
	 "encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type User struct {
	ID			string	`json:"id,omitempty"`
	Firstname	string	`json:"firstname,omitempty"`
	Lastname	string	`json:"lastname,omitempty"`
	Address		*Address `json:"address,omitempty"`
}

type Address struct {
	City	string	`json:"city,omitempty"`
	State	string	`json:"state,omitempty"`
}

var users []User 


func GetUsers(w http.ResponseWriter, r *http.Request) { 
	json.NewEncoder(w).Encode(users)
}

func GetUser(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&User{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = params["id"]
	users = append(users, user)
	json.NewEncoder(w).Encode(users)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) { 
	params := mux.Vars(r)
	for index, item := range users {
		if item.ID == params["id"] {
			users = append(users[:index], users[index+1:]...)
			break
		}
	json.NewEncoder(w).Encode(users)
	}
}


func main() {
	router := mux.NewRouter()

	// Creating Users
	users = append(users, User{ID: "1", Firstname: "Bjarne", Lastname:
	"Stroustrup", Address: &Address{City: "Sao Paulo", State: "Sao Paulo"}})

	users = append(users, User{ID: "2", Firstname: "Leonardo", Lastname:
	"Lima", Address: &Address{City: "Curitiba", State: "Parana"}})

	users = append(users, User{ID: "3", Firstname: "Scott", Lastname:
	"Pilgrin"})

	// Routes
	router.HandleFunc("/user", GetUsers).Methods("GET")
	router.HandleFunc("/user/{id}", GetUser).Methods("GET")
	router.HandleFunc("/user/{id}", CreateUser).Methods("POST")
	router.HandleFunc("/user/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
