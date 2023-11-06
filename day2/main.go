package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO
/*
  Write function for scoring whether you threw Rock, Paper, Sissors (Switch Statement?)
  Write a Function for scoring if you won or not (Switch Statement??)
  Need to keep track of total and return it in the end.
*/

func choiceScore(strat string) int {
	stratarray := strings.Split(strat, " ")
	score := 0
	switch stratarray[1] {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	}

	return score
}

/*
  A = Rock
  B = Paper
  C = Scissors
  X = Rock
  Y = Paper
  Z = Scissors
*/

func resultScore(strat string) int {
	score := 0
	switch strat {
	case "A X":
		score += 3
	case "A Y":
		score += 6
	case "B Y":
		score += 3
	case "B Z":
		score += 6
	case "C Z":
		score += 3
	case "C X":
		score += 6
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
		overallScore += choiceScore(fileScanner.Text())
		overallScore += resultScore(fileScanner.Text())
	}
	fmt.Println(overallScore)

}
