package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var forest [][]int8

func forestVisibility(forest [][]int8) int {
	visibleCount := 0
	// for the edges on the middle rows, will handle with others
	for i := 0; i <= len(forest)-1; i++ {
    fmt.Println("Count is : "+ fmt.Sprint(visibleCount))
		fmt.Printf("Starting on %v\n", forest[i])
		if i == 0 {
			visibleCount += len(forest[0])
		} else if i == len(forest)-1 {
			visibleCount += len(forest[len(forest)-1])
		} else {
			for j := 0; j <= len(forest[i])-1; j++ {
				currentValue := forest[i][j]
				currentIndex := j
				if j == 0 || j == len(forest[i])-1 {
          fmt.Printf("%v is added to the count 0/end\n", currentValue)
					visibleCount += 1
					continue
				}
				isVisible := true
				//left
				for left := j - 1; left >= 0; left-- {
					if currentValue <= forest[i][left] {
						isVisible = false
            break
					}
        }
				if isVisible {
          fmt.Printf("%v is added to the count, LEFT\n", currentValue)
					visibleCount += 1
					continue
				}
				isVisible = true
				//right
				for right := len(forest[i]) - 1; right > currentIndex; right-- {
					if currentValue <= forest[i][right] {
						isVisible = false
						break
					}
				}
				if isVisible {
          fmt.Printf("%v is added to the count, RIGHT\n", currentValue)
					visibleCount += 1
					continue
				}
				isVisible = true
				//up
				for up := i-1; up >= 0; up-- {
					if currentValue <= forest[up][currentIndex] {
						isVisible = false
						break
					}
				}
				if isVisible {
          fmt.Printf("%v is added to the count, UP\n", currentValue)
					visibleCount += 1
					continue
				}
				isVisible = true
				//down
				for down := len(forest) - 1; down > i; down-- {
					if currentValue <= forest[down][currentIndex] {
						isVisible = false
						break
					}
				}
				if isVisible {
          fmt.Printf("%v is added to the count, DOWN\n", currentValue)
					visibleCount += 1
					continue
				}
			}
		}
	}

	return visibleCount
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		arr := strings.Split(fileScanner.Text(), "")
		intArr := []int8{}
		for _, rune := range arr {
			intValue, _ := strconv.ParseInt(rune, 10, 8)
			intValue8 := int8(intValue)
			intArr = append(intArr, intValue8)
		}
		forest = append(forest, intArr)
	}
	fmt.Println(forestVisibility(forest))
}
