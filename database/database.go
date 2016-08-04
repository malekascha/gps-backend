package database

import (
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  // "fmt"
  "encoding/json"
  "os"
  "errors"
)

func StoreCoords (coords Point) error {
  user := os.Getenv("DB_USER")
  pass := os.Getenv("DB_PASS")

  mongo_uri := "mongodb://"+user+":"+pass+"@ds029745.mlab.com:29745/heroku_47clc7sm"

  if(len(coords.Message) == 0){
    return errors.New("must have message")
  }

  if(len(coords.Owner) == 0){
    return errors.New("must have valid owner")
  }

  if(coords.Geojson.Type != "Point"){
    return errors.New("must be GeoJSON of type Point")
  }

  if(len(coords.Geojson.Coordinates) != 2){
    return errors.New("must provide pair of valid coordinates")
  }

  session, err := mgo.Dial(mongo_uri)
  if(err != nil){
    return err
  }
  defer session.Close()

  c := session.DB("heroku_47clc7sm").C("coordinates")
  err = c.Insert(coords)

  return err
}

func RetrieveMessages (coords []float64, radius float64) ([]byte, error) {
  user := os.Getenv("DB_USER")
  pass := os.Getenv("DB_PASS")

  mongo_uri := "mongodb://"+user+":"+pass+"@ds029745.mlab.com:29745/heroku_47clc7sm"
  session, err := mgo.Dial(mongo_uri)
  if(err != nil){
    return []byte{}, err
  }
  defer session.Close()

  c := session.DB("heroku_47clc7sm").C("coordinates")
  var result []interface{}

  query := bson.M{"geojson": bson.M{"$near": bson.M{"$geometry": bson.M{"type": "Point", "coordinates": coords}, "$maxDistance": radius}}}
  err = c.Find(query).All(&result)
  if(err != nil){
    return []byte{}, err
  }

  jsonString, err := json.Marshal(result)
  if(err != nil){
    return []byte{}, err
  }
  return jsonString, nil
}

