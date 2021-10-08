package main

import (
	"appointy/controllers"
	"fmt"

	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {

	r := httprouter.New()
	uc := controllers.NewUserController(getSession())

	// All the url definitions are defined over here
	r.GET("/users/:id", uc.GetUser)
	r.POST("/users", uc.CreateUser)
	r.POST("/posts", uc.CreatePost)
	r.GET("/posts/:id", uc.GetPost)
	r.GET("/posts_user/:user_id", uc.GetPostFromUser) // In the task it was written to keep this url as /post/user/:id but that gives an error as /posts/:id is already being used

	fmt.Print("The api is being served at port:9000\n")

	http.ListenAndServe("localhost:9000", r) // start listner listens on port 9000
}

//function to initiate the mongodb connection
func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	return s
}
