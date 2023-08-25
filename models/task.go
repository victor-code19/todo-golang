package models

import "gopkg.in/mgo.v2/bson"

type Task struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Description string        `json:"description" bson:"description"`
}

type ViewTask struct {
	Id          string
	Description string
}
