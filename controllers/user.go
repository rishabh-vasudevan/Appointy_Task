package controllers

import (
	"appointy/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	Session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

var lock sync.Mutex

// This is the function that retrievs the data once the user Id has been provided
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") //extract the id from the get request

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) //return if error is found
	}

	oid := bson.ObjectIdHex(id) //converting id into form of bson.ObjectId

	u := models.Users{}

	//fiding the entry with same user_id
	if err := uc.Session.DB("appointy").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json") //sets up the http header so that it can carry response
	w.WriteHeader((http.StatusOK))                     //success return
	fmt.Fprintf(w, "%s\n", uj)                         //sends the object retrieved in the form of json

}

// This function is there to add the user into the Users collection in mongodb
func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	lock.Lock()
	u := models.Users{}

	json.NewDecoder(r.Body).Decode(&u) //decodes the body of the post request and stored it into u

	u.UserId = bson.NewObjectId() //creates a new bson.ObjectId for User

	/* Bcrypt is being used to add salt and hash to the password to make the password more secure
	and so that it cannot be reverse engineered */
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(u.Password), 8)

	u.Password = string(hashedPassword) // The password hash is converted to the string format to store it in the database

	uc.Session.DB("appointy").C("users").Insert(u) //inserts the object into the users collection

	uj, err := json.Marshal(u) //converts the object back into json

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) //returns the respose of created through http
	fmt.Fprintf(w, "%s\n", uj)        //returns the json file
}

// This function is there to create new entries into the posts collection in mongodb
func (uc UserController) CreatePost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	lock.Lock()
	u := models.Posts{}

	json.NewDecoder(r.Body).Decode(&u) //decode the body of the post request, decode it and store it in u

	u.PostId = bson.NewObjectId() // generate new bson Id
	u.PostTime = time.Now()       // time.Time() returns the current date time, that is added to the PostTime field

	uc.Session.DB("appointy").C("posts").Insert(u) //inserts data into the collection of posts in mongo db

	uj, err := json.Marshal(u) // converts it back into json file

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // sends successful creation message through http
	fmt.Fprintf(w, "%s\n", uj)        // sends the object back in the form of json
}

// This function is there to retrieve posts based on the post id
func (uc UserController) GetPost(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id") //retrieve the id from the get request

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id) // converts id to the form of bson.ObjectId

	u := models.Posts{} //create empty object of type posts

	// finds the post with the exact same post_id
	if err := uc.Session.DB("appointy").C("posts").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(u) // converts the data to json format
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	fmt.Fprintf(w, "%s\n", uj) // returns the found object in the form of json

}

// This funciton is there to retrieve all the posts made by a particular user
func (uc UserController) GetPostFromUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("user_id") //user id gets retrieved from the get request

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	result := []models.Posts{}

	/* Finds all the posts which were made by a particular user
	Pagination has also been used over here, the limit of retrieving posts has been set to 10  */
	if err := uc.Session.DB("appointy").C("posts").Find(nil).Select(bson.M{"user": id}).Limit(10).All(&result); err != nil {
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(result) //converts the data into the format of json
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader((http.StatusOK))
	fmt.Fprintf(w, "%s\n", uj) // returns all the objects that were sent by a particular user in the form of json

}
