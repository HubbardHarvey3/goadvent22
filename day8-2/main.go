package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkLeft(currentIndex int, currentValue int8, forest [][]int8, mainLoopIndex int) int8 {
  var numberOfTrees int8
  for i := currentIndex - 1; i >= 0; i-- {
    isDone := false
    switch {
      case currentValue > forest[mainLoopIndex][i] :
        numberOfTrees += 1
      case currentValue <= forest[mainLoopIndex][i] :
        numberOfTrees += 1
        isDone = true
    }
    if isDone {
      break
    }
  }
  return numberOfTrees
}

func checkRight(currentIndex int, currentValue int8, forest [][]int8, mainLoopIndex int) int8 {
  var numberOfTrees int8
  //fmt.Printf("In checkRight, working on %v\n", forest[mainLoopIndex])
  for i := currentIndex + 1; i < len(forest[mainLoopIndex]); i++ {
    isDone := false
    switch {
      case currentValue > forest[mainLoopIndex][i] :
        numberOfTrees += 1
      case currentValue <= forest[mainLoopIndex][i] :
        numberOfTrees += 1
        isDone = true
    }
    if isDone {
      break
    }
  }
  return numberOfTrees
}

func checkUp(currentIndex int, currentValue int8, forest [][]int8, mainLoopIndex int) int8 {
  var numberOfTrees int8
  //fmt.Printf("In checkUp, working on %v\n", forest[mainLoopIndex])
  for i := mainLoopIndex-1; i >= 0; i-- {
    isDone := false
    switch {
      case currentValue > forest[i][currentIndex] :
        numberOfTrees += 1
      case currentValue <= forest[i][currentIndex] :
        numberOfTrees += 1
        isDone = true
    }
    if isDone {
      break
    }
  }
  return numberOfTrees
}

func checkDown(currentIndex int, currentValue int8, forest [][]int8, mainLoopIndex int) int8 {
  var numberOfTrees int8
  //fmt.Printf("In checkDown, working on %v\n", forest[mainLoopIndex])
  for i := mainLoopIndex + 1; i <= len(forest)-1; i++ {
    isDone := false
    switch {
      case currentValue > forest[i][currentIndex] :
        numberOfTrees += 1
      case currentValue <= forest[i][currentIndex] :
        numberOfTrees += 1
        isDone = true
    }
    if isDone {
      break
    }
  }
  return numberOfTrees
}

var forest [][]int8

func forestVisibility(forest [][]int8) int64 {
  var visibleCount int64
	visibleCount = 0
	// for the edges on the middle rows, will handle with others
	for i := 0; i <= len(forest)-1; i++ {
		//fmt.Printf("Starting on %v\n", forest[i])
		if i == 0 {
      continue
		} else if i == len(forest)-1 {
      continue
		} else {
			for j := 0; j <= len(forest[i])-1; j++ {
				currentValue := forest[i][j]
				currentIndex := j
				if j == 0 || j == len(forest[i])-1 {
					continue
				}
        left := checkLeft(currentIndex, currentValue, forest, i)
        right := checkRight(currentIndex, currentValue, forest, i)
        up := checkUp(currentIndex, currentValue, forest, i)
        down := checkDown(currentIndex, currentValue, forest, i)
        total := int64(left) * int64(right) * int64(up) * int64(down)
        if visibleCount < int64(total) {
          visibleCount = int64(total)
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
