package main

import (
  "net/http"
)

func helloWorld (w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("Hello, world!"))
}