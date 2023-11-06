package main

import (
	"bufio"
	"fmt"
	"os"
)


func splitString(inputString string) ([]string, []string) {
  length := len(inputString)
  //fmt.Println(length)

  mid := length/2
  //fmt.Println(mid)
  var array1 []string
  var array2 []string
//  array1 := make([]string, mid)
//  array2 := make([]string, mid)

  for i :=0; i < length; i++ {
    if i <= mid-1 {
      array1 = append(array1, string(inputString[i]))
    } else {
      array2 = append(array2, string(inputString[i]))
    }
  }

  //fmt.Println(array1)
  //fmt.Println(array2)

  return array1, array2
}

func compareArray(array1, array2 []string) string{
  string := ""
  isFound := false
  for i:=0; i < len(array1); i++ {
    if isFound {
      break
    }
    for j:=0; j < len(array2); j++ {
      if isFound {
        break
      }
      if array1[i] == array2[j] {
        //fmt.Printf("FOUND ONE!! = %v\n", array1[i])
        string = array1[i]
        isFound = true
      }
    }
  }
  return string
}

var pointMap = map[string]int32{
  "a":1, "b":2, "c":3, "d":4, "e":5, "f":6, "g":7, "h":8, "i":9, "j":10, "k":11, "l":12, "m":13, "n":14, "o":15, "p":16, "q":17, "r":18, "s":19, "t":20, "u":21, "v":22, "w":23, "x":24, "y":25, "z":26,
  "A":27, "B":28, "C":29, "D":30, "E":31, "F":32, "G":33, "H":34, "I":35, "J":36, "K":37, "L":38, "M":39, "N":40, "O":41, "P":42, "Q":43, "R":44, "S":45, "T":46, "U":47, "V":48, "W":49, "X":50, "Y":51, "Z":52,
}
var score int32

func main() {

    readFile, err := os.Open("input.txt")

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)
    score=0
    for fileScanner.Scan() {
      str := compareArray(splitString(fileScanner.Text()))
      fmt.Println(str)
      score += pointMap[str]
      fmt.Println(str)
    }
    fmt.Println(score)
    readFile.Close()
}
