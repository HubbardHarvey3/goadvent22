package main

import (
  "testing"
)


func TestChoiceScoreAX(t *testing.T) {
  stringExample := "A X"
  want := 1
  score := choiceScore(stringExample)
  if want != score {
    t.Fatalf("choiceScore didn't return 1 for a case of A X")
  }


}

func TestChoiceScoreBY(t *testing.T) {
  stringExample := "B Y"
  want := 2

  score := choiceScore(stringExample)
  if want != score {
    t.Fatalf("choiceScore didn't return 2 for a case of B Y")
  }
}
