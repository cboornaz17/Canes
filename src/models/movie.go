package models

import (
  "gopkg.in/mgo.v2/bson"
)

// Represents a movie, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Movie struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Data        int64         `bson:"data" json:"data"`
}
