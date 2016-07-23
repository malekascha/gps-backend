package main

import (
  // "gopkg.in/mgo.v2"
  "net/http"
)

func main(){
  http.HandleFunc("/", helloWorld)
  http.ListenAndServe(":3000", http.DefaultServeMux)
}

