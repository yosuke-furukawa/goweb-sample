package tweets

import (
  "labix.org/v2/mgo"
  "os"
)

func getSession() *mgo.Session {
  uri := os.Getenv("MONGOLAB_URI")
  if uri == "" {
    uri = "localhost"
  }
  session, err := mgo.Dial(uri)
  if err != nil {
    panic(err)
  }
  return session
}

