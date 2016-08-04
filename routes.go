package main

import (
  "net/http"
  "github.com/malekascha/gps-backend/database"
  "encoding/json"
  "net/url"
  // "fmt"
  "strconv"
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

func getMessagesByRadius (w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, "Unsupported HTTP Method", http.StatusMethodNotAllowed)
    return
  }

  params, err := url.ParseQuery(r.URL.RawQuery)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  lat, err := strconv.ParseFloat(params.Get("latitude"), 64)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  long, err := strconv.ParseFloat(params.Get("longitude"), 64)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  radius, err := strconv.ParseFloat(params.Get("radius"), 64)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  messages, err := database.RetrieveMessagesByRadius([]float64{lat, long}, radius)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Write(messages)
}

func getMessagesByOwner (w http.ResponseWriter, r *http.Request) {
  if r.Method != "GET" {
    http.Error(w, "Unsupported HTTP Method", http.StatusMethodNotAllowed)
    return
  }

  params, err := url.ParseQuery(r.URL.RawQuery)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  owner := params.Get("owner")
  if(owner == ""){
    http.Error(w, "must provide name of owner", http.StatusBadRequest)
    return
  }

  messages, err := database.RetrieveMessagesByOwner(owner)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Write(messages)
}