package main

import (
  "testing"
)



func TestResultScoreAZ(t *testing.T) {
  testString := "A Z"
  want := 8
  score := resultScore(testString)
  if want != score {
    t.Fatalf("resultScore didn't return the correct value for case 'A Z'")
  }
}
