package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func stringToArray(input string) []string {
  value := strings.Split(input, "")
  return value
}


func compareAll(motherArray []string) int{
  start := time.Now()
  var score int
  var length int

  length = len(motherArray)

  for i := 0; i < length; i+=3 {
    array1 := stringToArray(motherArray[i])

      for _, v := range array1 {
        if (strings.Contains(motherArray[i+1], v) &&
          strings.Contains(motherArray[i+2], v)) {
            score +=pointMap[v]
            break
        }
      }
    }
  duration := time.Since(start)
  fmt.Printf("Total Execution time for CompareAll : %v\n", duration)
  fmt.Printf("Score = %v\n", score)
  return score
}

var pointMap = map[string]int{
  "a":1, "b":2, "c":3, "d":4, "e":5, "f":6, "g":7, "h":8, "i":9, "j":10, "k":11, "l":12, "m":13, "n":14, "o":15, "p":16, "q":17, "r":18, "s":19, "t":20, "u":21, "v":22, "w":23, "x":24, "y":25, "z":26,
  "A":27, "B":28, "C":29, "D":30, "E":31, "F":32, "G":33, "H":34, "I":35, "J":36, "K":37, "L":38, "M":39, "N":40, "O":41, "P":42, "Q":43, "R":44, "S":45, "T":46, "U":47, "V":48, "W":49, "X":50, "Y":51, "Z":52,
}
var score int

func main() {

    readFile, err := os.Open("input.txt")

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)
    var score int
    var motherArray []string
    for fileScanner.Scan() {
      motherArray = append(motherArray, fileScanner.Text())
    }

    score = compareAll(motherArray)

    fmt.Println(score)
    readFile.Close()
}
