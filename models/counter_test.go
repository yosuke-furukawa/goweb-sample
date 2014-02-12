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
  seq3, _ := Prev()
  seq4, _ := Prev()
  if seq != seq3  {
    t.Errorf("%d !=  %d", seq, seq3)
  }
  if seq4 != seq-1  {
    t.Errorf("%d !=  %d", seq, seq3)
  }
}


