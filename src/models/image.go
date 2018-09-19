package models

import (
  "gopkg.in/mgo.v2/bson"
)

// Represents an image, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Image struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
  Data        string        `bson:"data"        json:"data"`
}
