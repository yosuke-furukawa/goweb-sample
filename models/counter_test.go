package tweets

import (
  "testing"
  _ "fmt"
)

func TestNext(t *testing.T) {
  seq, _ := Next()
  seq2, _ := Next()
  if seq2 != seq + 1 {
    t.Errorf("%d !=  %d", seq, seq2)
  }
}


