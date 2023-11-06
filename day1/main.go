package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hi Ash")
	content, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("ERROR")
	}
	//fmt.Printf("%s", content)
	fmt.Println(bufio.ScanLines(content, true))
	//fmt.Println(reflect.TypeOf(content))
	for x, _ := range content {
		fmt.Println(x)
	}
}
