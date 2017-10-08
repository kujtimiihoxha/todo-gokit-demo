package io

import "gopkg.in/mgo.v2/bson"

type Todo struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Title    string        `json:"title"  bson:"title"`
	Complete bool          `json:"complete" bson:"complete"`
}
