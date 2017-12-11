package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	path := getInput()
	i, j, k, max := 0, 0, 0, 0
	for _, s := range path {
		switch s {
		case "n":
			j++
			k--
		case "ne":
			i++
			k--
		case "se":
			i++
			j--
		case "s":
			j--
			k++
		case "sw":
			i--
			k++
		case "nw":
			i--
			j++
		}
		d := dist(i, j, k)
		if d > max {
			max = d
		}
	}

	fmt.Println("Star 1: ", dist(i, j, k))
	fmt.Println("Star 2: ", max)
}

func dist(i, j, k int) int {
	return (abs(i) + abs(j) + abs(k)) / 2
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func getInput() []string {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(bytes), ",")
	return strs
}
