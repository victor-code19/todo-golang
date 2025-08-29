package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Task struct {
	Id          bson.ObjectID `json:"id" bson:"_id,omitempty"`
	Description string         `json:"description" bson:"description"`
}

type ViewTask struct {
    Id          string `json:"id"`
    Description string `json:"description"`
}