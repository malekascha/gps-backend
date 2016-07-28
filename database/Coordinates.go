package database

type Point struct {
  Message string `json:"message"`
  Coords coordinates `json:"coords"`
}

type coordinates struct {
  Lat int `json:"lat"`
  Long int `json:"long"`
}