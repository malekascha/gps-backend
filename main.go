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
  http.HandleFunc("/addcoords", addCoords)
  http.HandleFunc("/getmessages", getMessages)
  fmt.Println("Now listening at "+port)
  http.ListenAndServe(":"+port, http.DefaultServeMux)

  // r := database.RetrieveCoords([]int{25,25}, 300000000000)
  // fmt.Println(r)
}

