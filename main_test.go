package main

import (
	"appointy/controllers"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Amock struct {
}

func (a *Amock) TestGetUser(t *testing.T) {

	uc := controllers.NewUserController(getSession())
	req, err := http.NewRequest("GET", "http://localhost:9000/users/6160e3dfefd7173eb8292320", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(uc.GetUser)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
