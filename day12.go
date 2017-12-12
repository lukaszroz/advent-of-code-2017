package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := getLines()

	list := make([][]int, len(lines))
	reNumber := regexp.MustCompile("[0-9]+")
	for _, line := range lines {
		numbersString := reNumber.FindAllString(line, -1)
		numbers := atoi(numbersString)
		list[numbers[0]] = numbers[1:]
	}

	set := make(map[int]bool)
	addPipes(list, set, 0)

	fmt.Println("Start 1: ", len(set))

	groups := 1
	for len(set) < len(list) {
		for idx := range list {
			if !set[idx] {
				addPipes(list, set, idx)
				groups++
			}
		}
	}
	fmt.Println("Start 2: ", groups)
}

func addPipes(list [][]int, set map[int]bool, idx int) {
	set[idx] = true
	for _, n := range list[idx] {
		if !set[n] {
			addPipes(list, set, n)
		}
	}
}

func atoi(ss []string) []int {
	ii := make([]int, len(ss))
	var err error
	for idx, s := range ss {
		ii[idx], err = strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
	}
	return ii
}

func getLines() []string {
	bytes, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	strs := strings.Split(string(bytes), "\n")
	return strs
}
