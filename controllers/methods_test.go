package controllers

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
var client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
var uc = NewUserController(client)

func TestGetUser(t *testing.T) {

	req, err := http.NewRequest("GET", "http://localhost:9001/users/6161ced34c084e3f86b20a0e", nil)

	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	uc.GetUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

}

func CreateUser(t *testing.T) {

	var input = []byte(`{
		"name":"Rishabh",
		"email":"rishabhvasudevan19@gmail.com",
		"password":"12345"
	}`)

	req, err := http.NewRequest("POST", "http://localhost:9001/posts", bytes.NewBuffer(input))

	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	uc.GetUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

}

func GETPost(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:9001/posts/6161d3ba4ee6e84912f44251", nil)

	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	uc.GetUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
}

func CreatePost(t *testing.T) {
	var input = []byte(`{
		"user":"61619bdeefd71721b5fa90b3",
		"caption":"This is a photo",
		"image_url":"www.google.com"
	}`)

	req, err := http.NewRequest("POST", "http://localhost:9001/posts", bytes.NewBuffer(input))

	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	uc.GetUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

}

func GetPostFromUser(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:9001/posts/user/6161d3ba4ee6e84912f44251", nil)

	if err != nil {
		t.Fatal(err)
	}
	rec := httptest.NewRecorder()

	uc.GetUser(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.Status)
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}
}
