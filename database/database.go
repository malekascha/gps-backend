package database

import (
  "gopkg.in/mgo.v2"
)

func StoreCoords (coords Point) error {
  session, err := mgo.Dial("mongodb://admin:intelligenttraffic@ds029745.mlab.com:29745/heroku_47clc7sm")
  if(err != nil){
    panic(err)
  }
  defer session.Close()

  c := session.DB("heroku_47clc7sm").C("coordinates")
  err = c.Insert(coords)

  return err
}