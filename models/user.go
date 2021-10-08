package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// All the database structures are defined in this file

type Users struct {
	UserId   bson.ObjectId `json:"id" bson:"_id"`
	Name     string        `json:"name" bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"password" bson:"password"`
}

type Posts struct {
	PostId   bson.ObjectId `json:"id" bson:"_id"`
	User     string        `json:"user" bson:"user"`
	Caption  string        `json:"caption" bson:"caption"`
	ImageUrl string        `json:"image_url" bson:"image_url"`
	PostTime time.Time     `json:"post_time" bson:"post_time"`
}
