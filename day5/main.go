package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInstructions(text string) [3]int {
  textArr := strings.Split(text, " ")
  var intArr [3]int

  howMany, _ := strconv.Atoi(textArr[1])
  fromArray, _ := strconv.Atoi(textArr[3])
  toArray, _ := strconv.Atoi(textArr[5])

  fromArray -= 1
  toArray -= 1

  intArr[0] = howMany
  intArr[1] = fromArray
  intArr[2] = toArray
  fmt.Println(intArr)
  return intArr
}

func executeMoves(boxMoves int, toArr *[]string, fromArr *[]string) {
  for i:=0;i<boxMoves;i++ {
    //grab the end of fromArr and stick on end of toArr
    *toArr = append(*toArr, (*fromArr)[len(*fromArr)-1])
    //remove the recently moved box from the end of fromArr
    *fromArr = (*fromArr)[:len(*fromArr)-1]
  }
}


func main() {
    var motherArray [][]string
    motherArray = [][]string{
      {"F", "H", "B", "V", "R", "Q", "D", "P"},
      {"L", "D", "Z", "Q", "W", "V"},
      {"H", "L", "Z", "Q", "G", "R", "P", "C"},
      {"R", "D", "H", "F", "J", "V", "B"},
      {"Z", "W", "L", "C"},
      {"J", "R", "P", "N", "T", "G", "V", "M"},
      {"J", "R", "L", "V", "M", "B", "S"},
      {"D", "P", "J"},
      {"D", "C", "N", "W", "V"}}
    readFile, err := os.Open("input.txt")
    for i := range motherArray {
      fmt.Printf("Array # %v : %v\n",i+1, motherArray[i])
    }

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
      instructions := getInstructions(fileScanner.Text())
      executeMoves(instructions[0],
      &motherArray[instructions[2]],
      &motherArray[instructions[1]])
    }
    for i := range motherArray {
      fmt.Printf("Array # %v : %v\n",i+1, motherArray[i])
    }
    readFile.Close()
}

