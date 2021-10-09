package main

import (
	"appointy/controllers"
	"fmt"

	"net/http"

	"context"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Setup the connection with MongoDB
	r := mux.NewRouter()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	uc := controllers.NewUserController(client)


	// All the url definitions are defined over here
	r.HandleFunc("/users/{id}", uc.GetUser).Methods("GET")
	r.HandleFunc("/users", uc.CreateUser).Methods("POST")
	r.HandleFunc("/posts", uc.CreatePost).Methods("POST")
	r.HandleFunc("/posts/{id}", uc.GetPost).Methods("GET")
	r.HandleFunc("/posts/user/{id}", uc.GetPostFromUser).Methods("GET")

	fmt.Print("The api is being served at port:9001\n")

	http.ListenAndServe("localhost:9001", r) // start listner listens on port 9000
}
