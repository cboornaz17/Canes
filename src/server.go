package main

import (
	"encoding/json"
	"log"
	"net/http"
  "html/template"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	. "github.com/cboornaz17/pallas/src/config"
  . "github.com/cboornaz17/pallas/src/dao"
  . "github.com/cboornaz17/pallas/src/models"
)

var config = Config{}
var dao = ImagesDAO{}

// Serve index.html to handle user input
func IndexHandler(w http.ResponseWriter, r *http.Request) {
  tmpl, err := template.ParseFiles("../gui/index.html")

  if err != nil {
    // Log the detailed error
    log.Println(err.Error())
    // Return a generic "Internal Server Error" message
    http.Error(w, http.StatusText(500), 500)
    return
  }

  if err := tmpl.ExecuteTemplate(w, "index.html",  nil); err != nil {
    log.Println(err.Error())
    http.Error(w, http.StatusText(500), 500)
  }
}

func ConvertImage(w http.ResponseWriter, r *http.Request) {

}

// GET list of images
func AllImagesEndPoint(w http.ResponseWriter, r *http.Request) {
	images, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, images)
}

// GET a image by its ID
func FindImageEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	image, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Image ID")
		return
	}
	respondWithJson(w, http.StatusOK, image)
}

// POST a new image
func CreateImageEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var image Image
	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	image.ID = bson.NewObjectId()
	if err := dao.Insert(image); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, image)
}

// PUT update an existing image
func UpdateImageEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var image Image
	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(image); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing image
func DeleteImageEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var image Image
	if err := json.NewDecoder(r.Body).Decode(&image); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Delete(image); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
  r.HandleFunc("/", IndexHandler)
	r.HandleFunc("/images", CreateImageEndPoint).Methods("POST")
	r.HandleFunc("/images", AllImagesEndPoint).Methods("GET")
	//r.HandleFunc("/images", CreateImageEndPoint).Methods("POST")
	r.HandleFunc("/images", UpdateImageEndPoint).Methods("PUT")
	r.HandleFunc("/images", DeleteImageEndPoint).Methods("DELETE")
	r.HandleFunc("/images/{id}", FindImageEndpoint).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}
