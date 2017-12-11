package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input := getInput()

	score_total := 0
	score := 0
	count_total := 0
	var count int

	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '{':
			score++
			score_total += score
		case '}':
			score--
		case '<':
			i, count = skip(i+1, input)
			count_total += count
		}
	}

	fmt.Println("Star 1: ", score_total)
	fmt.Println("Star 2: ", count_total)
}

func skip(i int, input string) (int, int) {
	count := 0
	for ; i < len(input); i++ {
		switch input[i] {
		case '!':
			i++
		case '>':
			return i, count
		default:
			count++
		}
	}
	panic("> not found!")
}

func getInput() string {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
