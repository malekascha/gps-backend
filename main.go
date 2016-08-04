package main

import (
  "net/http"
  "os"
  "fmt"
  // "github.com/malekascha/gps-backend/database"
)

func main(){

  port := os.Getenv("PORT")

  if (port == "") {
    port = "3000"
  }

  http.HandleFunc("/", helloWorld)
  http.HandleFunc("/api/addmessage", addCoords)
  http.HandleFunc("/api/getmessagesbyradius", getMessages)
  fmt.Println("Now listening at "+port)
  http.ListenAndServe(":"+port, http.DefaultServeMux)
}

