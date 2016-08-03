package database

type Point struct {
  Message string `json:"message"`
  Geojson coordinates `json:"geojson"`
  Owner int `json:"owner"`
}

type coordinates struct {
  Type string `json:"type"`
  Coordinates []float64 `json:"coordinates"`
}