package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)



func main() {
    readFile, err := os.Open("input.txt")

    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)

    fileScanner.Split(bufio.ScanLines)

    for fileScanner.Scan() {
      fmt.Println(fileScanner.Text())
    }
    readFile.Close()
}

