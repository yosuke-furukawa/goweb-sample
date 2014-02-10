package tweets

import (
  "labix.org/v2/mgo"
  "os"
  "regexp"
  "fmt"
)

func getSessionAndDB() (*mgo.Session, *mgo.Database) {
  envuri := os.Getenv("MONGOLAB_URI")
  dburi := "localhost"
  dbname := "tweets"
  if envuri != "" {
    uriRegExp := regexp.MustCompile(`(.*)\/(\w+)`)
    connects := uriRegExp.FindStringSubmatch(envuri)
    dburi = connects[1]
    dbname = connects[2]
    fmt.Println(dburi)
    fmt.Println(dbname)
  }

  session, err := mgo.Dial(dburi)
  if err != nil {
    panic(err)
  }
  return session, session.DB(dbname)
}

