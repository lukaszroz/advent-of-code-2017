package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args[1]

	fmt.Println("1st star: ", nsum(1, list))
	fmt.Println("2nd star: ", nsum(len(list) / 2, list))
}

func nsum(n int, list string) int {
	sum := 0
	for i := range list {
		if(list[i] == list[(i + n) % len(list)]) {
			sum += int(list[i] - '0')
		}
	}
	return sum
}