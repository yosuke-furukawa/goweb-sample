package tweets

import (
  "testing"
  _ "fmt"
)

func Test_getSessionAndDB(t *testing.T) {
  session, database := getSessionAndDB()
  defer session.Close()
  if database == nil {
    t.Fatalf("database is nil")
  }
  if session == nil {
    t.Errorf("Session is nil")
  }
}



