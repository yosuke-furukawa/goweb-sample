package tweets

import (
  "labix.org/v2/mgo"
  "os"
  "regexp"
)

func getSessionAndDB() (*mgo.Session, *mgo.Database) {
  envuri := os.Getenv("MONGOLAB_URI")
  dburi := "localhost"
  dbname := "tweets"
  if envuri != "" {
    uriRegExp := regexp.MustCompile(`(.*)\/(\w+)`)
    connects := uriRegExp.FindStringSubmatch(envuri)
    dburi = connects[0]
    dbname = connects[2]
  }

  session, err := mgo.Dial(dburi)
  session.SetSafe(&mgo.Safe{})
  if err != nil {
    panic(err)
  }
  return session, session.DB(dbname)
}

