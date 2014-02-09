package tweets

import (
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type Counter struct {
  Id string "_id,omitempty"
  Seq int "seq"
}

func Init() error {
  session := getSession()
  defer session.Close()
  counters := session.DB("tweets").C("counters")
  c := &Counter{}
  c.Id = "tweet_id"
  c.Seq = 0
  err := counters.Insert(c)
  return err
}

func Next() (next int, err error) {
  session := getSession()
  defer session.Close()
  counters := session.DB("tweets").C("counters")

  change := mgo.Change{
    Update : bson.M{"$inc" : bson.M{"seq" : 1}},
    ReturnNew : true,
  }
  c := &Counter{}
  _, err = counters.FindId("tweet_id").Apply(change, &c)
  return c.Seq, err
}
