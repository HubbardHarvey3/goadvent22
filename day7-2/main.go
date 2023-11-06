package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Represents a directory or file in the filesystem
type FileSystemEntry struct {
	Name     string
	IsParent bool
	Size     int
}

func main() {
	readFile, err := os.Open("input.txt")
	//rootDir := FileSystemEntry{Name: "/", IsFile: false, Size: 0}
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	//parentDirs := map[string]int{"/":0}
	parentDirs := []FileSystemEntry{}
	for fileScanner.Scan() {
		//split the command up by space
		command := strings.Split(fileScanner.Text(), " ")
		if command[0] != "$" {
			//convert string to Int
			sizeInt, _ := strconv.Atoi(command[0])
			for i := range parentDirs {
				if parentDirs[i].IsParent {
					parentDirs[i].Size = parentDirs[i].Size + sizeInt
				}
			}
		} else if command[1] == "cd" && command[2] != ".." {
			entry := FileSystemEntry{Name: command[2], IsParent: true, Size: 0}
			parentDirs = append(parentDirs, entry)
		} else if command[1] == "cd" && command[2] == ".." {
			for i := len(parentDirs) - 1; i >= 0; i-- {
				if parentDirs[i].IsParent {
					parentDirs[i].IsParent = false
					break
				}
			}
		}
	}
  fmt.Println(parentDirs[0])
  emptySpace := 70000000 - parentDirs[0].Size
  spaceNeeded := 30000000 - emptySpace
	var finalSlice []FileSystemEntry
	for _, entry := range parentDirs {
		if entry.Size >= spaceNeeded {
      finalSlice = append(finalSlice, entry)
		}
	}
  fmt.Println(finalSlice)
  bestChoice := finalSlice[0]
  for i := range finalSlice {
    if (finalSlice[i].Size < bestChoice.Size) {
      bestChoice = finalSlice[i]
    }
  }
  fmt.Printf("Directory %v is the best with a size of %d\n", bestChoice.Name, bestChoice.Size)
}
