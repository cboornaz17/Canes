package dao

import (
	"log"

	. "github.com/cboornaz17/pallas/src/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ImagesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "images"
)

// Establish a connection to database
func (imgDao *ImagesDAO) Connect() {
	session, err := mgo.Dial(imgDao.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(imgDao.Database)
}

// Find list of images
func (imgDao *ImagesDAO) FindAll() ([]Image, error) {
	var images []Image
	err := db.C(COLLECTION).Find(bson.M{}).All(&images)
	return images, err
}

// Find a image by its id
func (imgDao *ImagesDAO) FindById(id string) (Image, error) {
	var image Image
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&image)
	return image, err
}

// Insert a image into database
func (imgDao *ImagesDAO) Insert(image Image) error {
	err := db.C(COLLECTION).Insert(&image)
	return err
}

// Delete an image from our database
func (imgDao *ImagesDAO) Delete(image Image) error {
	err := db.C(COLLECTION).Remove(&image)
	return err
}

// Update an existing image
func (imgDao *ImagesDAO) Update(image Image) error {
	err := db.C(COLLECTION).UpdateId(image.ID, &image)
	return err
}
