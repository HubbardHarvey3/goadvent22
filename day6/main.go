package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var buffer []string

func checkForDuplicates(strArr []string) bool {
	// return false if duplicates found
	// return true if buffer is unique
	for i := 0; i < len(strArr); i++ {
    for j:=0; j<len(strArr);j++{
      if j != i {
        if strArr[i] == strArr[j] {
          return false
        }
      }
    }
  }

  return true
}

func crawlCheck(strArr []string) int {
	count := 0
	buffer = strArr[:14]
	for i := 14; i < len(strArr); i++ {
    if checkForDuplicates(buffer) {
      fmt.Println("Possible Win?")
      return count + i
    } else {
      fmt.Println("Keep on looking")
    }
    //remove the first element
    buffer = buffer[1:14]
    //add new to end of slice
    buffer = append(buffer, strArr[i])
	}
	return 0
}

func main() {
	readFile, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	var theString []string
  var result int
	for fileScanner.Scan() {
		theString = strings.Split(fileScanner.Text(), "")
    result = crawlCheck(theString)
	}
  fmt.Println(result)
	readFile.Close()
}
