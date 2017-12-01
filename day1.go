package main

import (
	"fmt"
	"os"
)

func main() {
	list := os.Args[1]
	sum := 0
	for i := range list {
		if(list[i] == list[(i + 1) % len(list)]) {
			sum += int(list[i] - '0')
		}
	}
	fmt.Println(sum)
}
