package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pubsubexample/models"
	"pubsubexample/pubsub"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/users", handleUserRequest)
	http.HandleFunc("/add-user", handleUserRequest)
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func handleUserRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		handleUserPost(r)
	}
	w.WriteHeader(201)
}

func handleUserPost(request *http.Request) {
	var user models.User

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&user)
	fmt.Printf("Received POST %v", user)
	if err != nil {
		panic(err)
	}

	client := pubsub.New("go-pubsub-quickstart", "PubSubTestTopic")

	defer client.Close()

	subsErr := client.Subscribe("go-pubsub-quickstart", "PubSubTestTopic-sub")
	if subsErr != nil {
		log.Fatalf("Failed to subscribe %v", subsErr)
	}

	encodedMessage, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("could not encode message %v", err)
	}

	message := client.Message(encodedMessage)
	publishErr := client.Publish("go-pubsub-quickstart", client.GetTopicID(), message)

	if publishErr != nil {
		log.Fatalf("Failed to publish")
	}

	fmt.Printf("Execution finished.\n")
}
