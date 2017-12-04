package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")

	star1 := 0
	star2 := 0
	for _, line := range lines {
		strs := strings.Split(string(line), " ")
		star1 += hasNoDupes(strs)
		star2 += hasNoAnagrams(strs)
	}
	fmt.Println("1st star: ", star1)
	fmt.Println("2nd star: ", star2)
}

func hasNoAnagrams(strs []string) int {
	for idx, word := range strs {
		runes := []rune(word)
		sort.Sort(sortRunes(runes))
		strs[idx] = string(runes)
	}
	return hasNoDupes(strs)
}

type sortRunes []rune

func (a sortRunes) Len() int           { return len(a) }
func (a sortRunes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sortRunes) Less(i, j int) bool { return a[i] < a[j] }

func hasNoDupes(strs []string) int {
	for idx, word := range strs {
		for j := idx + 1; j < len(strs); j++ {
			if word == strs[j] {
				return 0
			}
		}
	}
	return 1
}
