package database

type Point struct {
  Message string `json:"message"`
  Geojson coordinates `json:"geojson"`
}

type coordinates struct {
  Type string `json:"type"`
  Coordinates []float64 `json:"coords"`
}