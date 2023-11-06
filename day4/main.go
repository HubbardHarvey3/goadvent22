package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
  "strconv"
)

func compareSections(text string) bool {
  var isSame bool
  // split the string up
  testArr := strings.Split(text, ",")

  // split it again by the -
  firstSection := strings.Split(testArr[0], "-")
  secondSection := strings.Split(testArr[1], "-")

  firstMin, _ := strconv.Atoi(firstSection[0])
  firstMax , _:= strconv.Atoi(firstSection[1])
  secondMin , _:= strconv.Atoi(secondSection[0])
  secondMax , _:= strconv.Atoi(secondSection[1])

  // compare the numbers to determine if the sections are the same
  if (firstMin <= secondMin && firstMax >= secondMax) ||
     (firstMin >= secondMin && firstMax <= secondMax) {
      fmt.Println("Found an inside site")
      fmt.Printf("%d-%d\n",firstMin, firstMax)
      fmt.Printf("%d-%d\n",secondMin, secondMax)
      isSame = true
  }

  return isSame
}

func main() {

    readFile, err := os.Open("input.txt")

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    count := 0

    for fileScanner.Scan() {
      if compareSections(fileScanner.Text()) {
        count++
      }
    }
    fmt.Printf("Count of duplicate sections = %v\n", count)
    readFile.Close()
}

