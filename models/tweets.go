package tweets

import (
  "labix.org/v2/mgo/bson"
  "fmt"
  "time"
)

type Tweet struct {
  Id int `bson:"_id,omitempty"  json:"_id"`
  Body string `bson:"body"  json:"body"`
  UpdatedAt int `bson:"updatedAt"  json:"updatedAt"`
}

func (t *Tweet) Insert() error {
  session, db := getSessionAndDB()
  defer session.Close()
  tweets := db.C("tweets")
  t.Id, _ = Next()
  t.UpdatedAt = int(time.Now().Unix())
  fmt.Println(t.Id)
  err := tweets.Insert(t)
  fmt.Println(err)
  return err
}

func FindById(id int) (t *Tweet, err error) {
  session, db := getSessionAndDB()
  defer session.Close()
  tweets := db.C("tweets")
  t = &Tweet{}
  err = tweets.FindId(id).One(&t)
  return t, err
}

func FindTweets(lastId int) (results []Tweet, err error) {
  if lastId == 0 {
    lastId = 1000000
  }
  session, db := getSessionAndDB()
  defer session.Close()
  tweets := db.C("tweets")
  err = tweets.Find(bson.M{"_id" : bson.M{"$lt" : lastId}}).Limit(10).Sort("-_id").All(&results)
  return results, err
}

func DeleteById(id int) error {
  session, db := getSessionAndDB()
  defer session.Close()
  tweets := db.C("tweets")
  err := tweets.RemoveId(id)
  return err
}

func (t *Tweet) Delete() error {
  session, db := getSessionAndDB()
  defer session.Close()
  tweets := db.C("tweets")
  err := tweets.RemoveId(t.Id)
  return err
}
