package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	lines := getLines()

	input := make(map[int]int, len(lines))
	reNumber := regexp.MustCompile("[0-9]+")
	for _, line := range lines {
		numbersString := reNumber.FindAllString(line, -1)
		numbers := atoi(numbersString)
		input[numbers[0]] = numbers[1]
	}

	fmt.Println("Start 1: ", severity(input, 0))
	fmt.Println("Start 2: ", findDelay(input))
}

type tuple struct {
	key, period int
}

type byPeriod []tuple

func (t byPeriod) Len() int           { return len(t) }
func (t byPeriod) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t byPeriod) Less(i, j int) bool { return t[i].period < t[j].period }

func findDelay(input map[int]int) int {
	tuples := make([]tuple, len(input))
	i := 0
	for key, val := range input {
		tuples[i] = tuple{key, (val - 1) * 2}
		i++
	}
	sort.Sort(byPeriod(tuples))
	for delay := 0; ; delay++ {
		if !caught(tuples, delay) {
			return delay
		}
	}
}

func caught(input []tuple, delay int) bool {
	for _, tup := range input {
		if (tup.key+delay)%(tup.period) == 0 {
			return true
		}
	}
	return false
}

func severity(input map[int]int, delay int) int {
	sum := 0
	for key, val := range input {
		if (key+delay)%((val-1)*2) == 0 {
			sum += key * val
		}
	}
	return sum
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
