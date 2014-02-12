package tweets

import (
  "testing"
  "fmt"
)

func TestInsert(t *testing.T) {
  tweet := &Tweet{}
  tweet.Body = "test"
  tweet.Insert()
  
  tweet2, _ := FindById(tweet.Id)
  if tweet2.Id != tweet.Id {
    t.Errorf("ID is different. %s, %s", tweet.Id, tweet2.Id)
  }
  if tweet2.Body != tweet.Body {
    t.Errorf("Body is different.")
  }
  if tweet2.UpdatedAt != tweet.UpdatedAt {
    t.Errorf("UpdatedAt is different.")
  }
  tweet.Delete()
}

func TestFindAll(t *testing.T) {
  results, _ := FindTweets(0)
  fmt.Println(results)
}

func TestNotfountTweet(t *testing.T) {
  results, _ := FindTweets(-1)
  fmt.Println(results)
}

