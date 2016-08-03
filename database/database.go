package database

import (
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  // "fmt"
  "encoding/json"
)

func StoreCoords (coords Point) error {
  session, err := mgo.Dial("mongodb://admin:intelligenttraffic@ds029745.mlab.com:29745/heroku_47clc7sm")
  if(err != nil){
    return err
  }
  defer session.Close()

  c := session.DB("heroku_47clc7sm").C("coordinates")
  err = c.Insert(coords)

  return err
}

func RetrieveMessages (coords []float64, radius float64) ([]byte, error) {
  session, err := mgo.Dial("mongodb://admin:intelligenttraffic@ds029745.mlab.com:29745/heroku_47clc7sm")
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