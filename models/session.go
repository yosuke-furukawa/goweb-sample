package tweets

import (
  "labix.org/v2/mgo"
)

func getSession() *mgo.Session {
  session, err := mgo.Dial("localhost")
  if err != nil {
    panic(err)
  }
  return session
}

