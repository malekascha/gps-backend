package main

import (
  "net/http"
  "os"
  "fmt"
)

func main(){

  port := os.Getenv("PORT")

  if (port == "") {
    port = "3000"
  }

  http.HandleFunc("/", helloWorld)
  http.HandleFunc("/addcoords", addCoords)
  fmt.Println("Now listening at "+port)
  http.ListenAndServe(":"+port, http.DefaultServeMux)
}

