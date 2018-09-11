// The client.go file serves up a templated webpage on a local port, receives an image file from a user, and passes that file onto the server.go program
package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  index, err := ioutil.ReadFile(string("../../gui/index.html"))
  if err != nil {
    w.WriteHeader(404)
    w.Write([]byte("404 Something went wrong - " + http.StatusText(404)))
  } else {
    w.Write(index)
  }
}

func imageReceiver(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Received image")
  indexHandler(w, r)
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/getImage", imageReceiver)
  http.ListenAndServe(":3000", nil)
  fmt.Println("this program will eventually send images from a client to a server")
}
