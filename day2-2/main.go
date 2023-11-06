package main

import (
	"bufio"
	"fmt"
	"os"
)


/*
  A = Rock +1
  B = Paper +2
  C = Scissors + 3
  X = Lose 0
  Y = Draw 3
  Z = Win 6
*/

func resultScore(strat string) int {
	score := 0
	switch strat {
	case "A X":
		score += 3
	case "A Y":
		score += 4
	case "A Z":
		score += 8
	case "B X":
		score += 1
	case "B Y":
		score += 5
	case "B Z":
		score += 9
  case "C X":
    score += 2
  case "C Y":
    score += 6
  case "C Z":
    score += 7
	}

	return score
}

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)
	overallScore := 0
	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
		overallScore += resultScore(fileScanner.Text())
	}
	fmt.Println(overallScore)

}
