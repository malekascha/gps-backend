package main

import (
  "net/http"
  "github.com/malekascha/gps-backend/database"
  "encoding/json"
)

func helloWorld (w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, world!"))
}

func addCoords (w http.ResponseWriter, r *http.Request) {
  if r.Method != "POST" {
    http.Error(w, "Unsupported HTTP Method", http.StatusMethodNotAllowed)
    return
  }

  point := database.Point{}
  err := json.NewDecoder(r.Body).Decode(&point)

  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  err = database.StoreCoords(point)
  
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Write([]byte("Point added to database"))
}

