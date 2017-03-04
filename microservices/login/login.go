package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	Level    string
	Username string
	password string
}

type LoginReq struct {
	Username string
	Password string
}

var admin = user{Level: "1", Username: "Admin", password: "admin"}
var paul = user{Level: "2", Username: "Stern", password: "mrpaulstern"}
var guest = user{Level: "3", Username: "Guest", password: "password"}

var users = map[string]user{admin.Username: admin, paul.Username: paul, guest.Username: guest}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "LOGIN")
	var u LoginReq
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	person, found := users[u.Username]
	if !found {
		http.Error(w, "Not found", 404)
		return
	}

	if u.Password == person.password {
		w.Header().Set("Content-Type", "application/json")
		jsonData, _ := json.Marshal(users[person.Username])
		w.Write(jsonData)
		fmt.Println("Logged in user", u)
	} else {
		http.Error(w, "Not authorized", 401)
		fmt.Println("Bad Login request", u)
	}

}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}
