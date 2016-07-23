package main

import (
  // "gopkg.in/mgo.v2"
  "net/http"
  "os"
)

func main(){

  port := os.Getenv("PORT")

  if (port == "") {
    port = ":3000"
  }

  http.HandleFunc("/", helloWorld)
  http.ListenAndServe(port, http.DefaultServeMux)
}

