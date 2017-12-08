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

	regs := make(map[string]int)
	//a dec -511 if x >= -4
	re := regexp.MustCompile("([a-z]+) (dec|inc) (-?[0-9]+) if ([a-z]+) (..?) (-?[0-9]+)")
	var processMax int = -1 << 31
	for _, line := range lines {
		tokens := re.FindAllStringSubmatch(line, -1)
		if condition(tokens[0][4:], regs) {
			v := operation(tokens[0][1:4], regs)
			if v > processMax {
				processMax = v
			}
		}
	}

	var completedMax int = -1 << 31
	for _, v := range regs {
		if v > completedMax {
			completedMax = v
		}
	}

	fmt.Println("Star 1: ", completedMax)
	fmt.Println("Star 2: ", processMax)
}

func getLines() []string {
	file, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	return strings.Split(string(file), "\n")
}

func condition(tokens []string, regs map[string]int) bool {
	n := atoi(tokens[2])
	r := regs[tokens[0]]
	switch tokens[1] {
	case ">":
		return r > n
	case ">=":
		return r >= n
	case "<":
		return r < n
	case "<=":
		return r <= n
	case "==":
		return r == n
	case "!=":
		return r != n
	}
	panic("Parsing failed!")
}

func operation(tokens []string, regs map[string]int) int {
	n := atoi(tokens[2])
	switch tokens[1] {
	case "inc":
		regs[tokens[0]] += n
	case "dec":
		regs[tokens[0]] -= n
	}
	return regs[tokens[0]]
}

func atoi(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}
