package models

import (
  "gopkg.in/mgo.v2/bson"
)

// Represents an image, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Image struct {
	ID          bson.ObjectId `bson:"_id"    json:"id"`
	Name        string        `bson:"name"   json:"name"`
  Data        string        `bson:"data"   json:"data"`
}
