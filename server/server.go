package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pubsubexample/models"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/users", handleUserRequest)
	http.HandleFunc("/add-user", handleUserRequest)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handleUserPost(w, r)
	}
	w.WriteHeader(201)
}

func handleUserPost(w http.ResponseWriter, r *http.Request) {
	var user models.User

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(201)

	err = models.SaveUser(user.Username, user.Name)
	if err != nil {
		panic(err)
	}
}
